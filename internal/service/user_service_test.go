package service

import (
	"context"
	"testing"

	"go-ecommerce-backend-api/internal/utils/sendto"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

func TestGetUserByID_Found(t *testing.T) {
	redisClient := redis.NewClient(&redis.Options{Addr: "localhost:6379", DB: 0})
	emailSender := sendto.NewMandrillEmailSender(sendto.SMTPConfig{})
	svc := NewUserService(redisClient, emailSender)

	user, err := svc.GetUserByID(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, "Alice Nguyen", user.Username)
	require.Equal(t, "Alice Nguyen", user.Username)
}

func TestGetUserByID_NotFound(t *testing.T) {
	redisClient := redis.NewClient(&redis.Options{Addr: "localhost:6379", DB: 0})
	emailSender := sendto.NewMandrillEmailSender(sendto.SMTPConfig{})
	svc := NewUserService(redisClient, emailSender)

	user, err := svc.GetUserByID(context.Background(), "999")

	require.Error(t, err)
	require.Nil(t, user)
}
