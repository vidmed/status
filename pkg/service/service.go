package service

import (
	"context"

	"github.com/vidmed/status/pkg/service/dto"
	"github.com/vidmed/status/pkg/service/storage"
)

// Service describes the status service.
type Service interface {
	GetStatus(ctx context.Context, req *dto.OrderStatusRequest) (*dto.OrderStatus, error)
	GetStatuses(ctx context.Context, req *dto.OrderStatusesRequest) (*dto.OrderStatuses, error)
	GetStatusHistory(ctx context.Context, req *dto.StatusHistoryRequest) (*dto.StatusHistory, error)
}

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(Service) Service

// New creates and returns a new status service instance
func New(storage storage.Storage) Service {
	return &service{
		storage: storage,
	}
}

// NewWithMiddleware returns a new status service with all of the expected middlewares wired in.
func NewWithMiddleware(storage storage.Storage, middleware []Middleware) Service {
	svc := New(storage)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

// service implements the status Service
type service struct {
	storage storage.Storage
}

func (s service) GetStatus(ctx context.Context, req *dto.OrderStatusRequest) (*dto.OrderStatus, error) {
	return s.storage.GetStatus(ctx, req)
}

func (s service) GetStatuses(ctx context.Context, req *dto.OrderStatusesRequest) (*dto.OrderStatuses, error) {
	return s.storage.GetStatuses(ctx, req)
}

func (s service) GetStatusHistory(ctx context.Context, req *dto.StatusHistoryRequest) (*dto.StatusHistory, error) {
	return s.storage.GetStatusHistory(ctx, req)
}
