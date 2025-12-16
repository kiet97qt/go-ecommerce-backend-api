package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetUserByID_Found(t *testing.T) {
	svc := NewUserService()

	user, err := svc.GetUserByID(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, "1", user.ID)
	require.Equal(t, "Alice Nguyen", user.Username)
}

func TestGetUserByID_NotFound(t *testing.T) {
	svc := NewUserService()

	user, err := svc.GetUserByID(context.Background(), "999")

	require.Error(t, err)
	require.Nil(t, user)
}
