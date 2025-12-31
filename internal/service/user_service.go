package service

import (
	"context"
	"fmt"
	"time"

	"go-ecommerce-backend-api/internal/models"
	"go-ecommerce-backend-api/internal/utils/crypto"
	"go-ecommerce-backend-api/internal/utils/random"
	"go-ecommerce-backend-api/internal/utils/sendto"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// IUserService exposes operations for user resources.
type IUserService interface {
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	RegisterUser(ctx context.Context, email string) error
}

type userService struct {
	users       map[string]models.User
	redisClient *redis.Client
	emailSender sendto.EmailSender
	otpTTL      time.Duration
}

// NewUserService returns a mock user service populated with sample data.
func NewUserService(redisClient *redis.Client, emailSender sendto.EmailSender) IUserService {
	return &userService{
		users: map[string]models.User{
			"1": {ID: uuid.New(), Username: "Alice Nguyen"},
			"2": {ID: uuid.New(), Username: "Bob Tran"},
		},
		redisClient: redisClient,
		emailSender: emailSender,
		otpTTL:      5 * time.Minute,
	}
}

func (s *userService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	user, ok := s.users[id]
	if !ok {
		return nil, fmt.Errorf("user %s not found", id)
	}
	return &user, nil
}

// RegisterUser tạo OTP và gửi email, đồng thời lưu OTP vào Redis.
func (s *userService) RegisterUser(ctx context.Context, email string) error {
	if s.redisClient == nil || s.emailSender == nil {
		return fmt.Errorf("service not properly initialized")
	}

	// hashEmail
	emailHash := crypto.HashEmail(email)
	key := fmt.Sprintf("otp:%s", emailHash)

	// avoid user spam... trong TTL: nếu key còn tồn tại, từ chối gửi lại
	exists, err := s.redisClient.Exists(ctx, key).Result()
	if err == nil && exists > 0 {
		return fmt.Errorf("otp already sent recently, please try again later")
	}

	// new OTP và save OTP in redis với expiration time
	otp, err := random.NumericOTP(6)
	if err != nil {
		return fmt.Errorf("failed to generate otp: %w", err)
	}

	if err := s.redisClient.Set(ctx, key, otp, s.otpTTL).Err(); err != nil {
		return fmt.Errorf("failed to save otp: %w", err)
	}

	// send email với OTP
	if err := s.emailSender.SendOTP(ctx, email, otp); err != nil {
		return fmt.Errorf("failed to send otp email: %w", err)
	}

	return nil
}
