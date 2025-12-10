package service

import "context"

// PingService exposes the interface for handling ping requests.
type PingService interface {
	Ping(ctx context.Context) string
}

type pingService struct{}

// NewPingService creates a new PingService implementation.
func NewPingService() PingService {
	return &pingService{}
}

func (pingService) Ping(ctx context.Context) string {
	return "pong"
}
