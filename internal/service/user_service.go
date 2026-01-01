package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/models"
	db "go-ecommerce-backend-api/internal/models/db"
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
	CreateUserWithOTP(ctx context.Context, email, otp string) (*db.User, error)
}

type userService struct {
	users       map[string]models.User
	redisClient *redis.Client
	emailSender sendto.EmailSender
	otpTTL      time.Duration
	dbQueries   *db.Queries
}

// NewUserService returns a mock user service populated with sample data.
func NewUserService(redisClient *redis.Client, emailSender sendto.EmailSender) IUserService {
	var queries *db.Queries
	if global.DB != nil {
		if sqlDB, err := global.DB.DB(); err == nil {
			queries = db.New(sqlDB)
		}
	}

	return &userService{
		users: map[string]models.User{
			"1": {ID: uuid.New(), Username: "Alice Nguyen"},
			"2": {ID: uuid.New(), Username: "Bob Tran"},
		},
		redisClient: redisClient,
		emailSender: emailSender,
		otpTTL:      5 * time.Minute,
		dbQueries:   queries,
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

// CreateUserWithOTP xác thực OTP và tạo user mới trong DB nếu email chưa tồn tại.
// Ở đây dùng cột username làm email.
func (s *userService) CreateUserWithOTP(ctx context.Context, email, otp string) (*db.User, error) {
	if s.redisClient == nil {
		return nil, fmt.Errorf("redis client not initialized")
	}
	if s.dbQueries == nil {
		return nil, fmt.Errorf("database not initialized")
	}

	emailHash := crypto.HashEmail(email)
	key := fmt.Sprintf("otp:%s", emailHash)

	// Lấy OTP từ Redis
	savedOTP, err := s.redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("otp expired or not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get otp: %w", err)
	}
	if savedOTP != otp {
		return nil, fmt.Errorf("invalid otp")
	}

	// Kiểm tra user đã tồn tại chưa (username = email)
	if _, err := s.dbQueries.GetUserByUsername(ctx, email); err == nil {
		return nil, fmt.Errorf("user already exists")
	} else if err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}

	now := time.Now().Unix()
	newID := uuid.New().String()

	if err := s.dbQueries.CreateUser(ctx, db.CreateUserParams{
		ID:        newID,
		Username:  email,
		IsActive:  true,
		CreatedAt: now,
		UpdatedAt: now,
	}); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Xoá OTP sau khi dùng
	_ = s.redisClient.Del(ctx, key).Err()

	created, err := s.dbQueries.GetUserByID(ctx, newID)
	if err != nil {
		return nil, fmt.Errorf("failed to load created user: %w", err)
	}

	return &created, nil
}
