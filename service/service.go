package service

import (
	"context"

	"github.com/vidmed/status/dto"
	"github.com/vidmed/status/storage"
)

// Service describes the status service.
type Service interface {
	GetStatus(ctx context.Context, req *dto.OrderStatusRequest) (*dto.OrderStatus, error)
	GetStatuses(ctx context.Context, req *dto.OrderStatusesRequest) (*dto.OrderStatuses, error)
	GetStatusHistory(ctx context.Context, req *dto.StatusHistoryRequest) (*dto.StatusHistory, error)
}

// service implements the status Service
type service struct {
	storage storage.Storage
}

// NewService creates and returns a new status service instance
func NewService(storage storage.Storage) Service {
	return &service{
		storage: storage,
	}
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
