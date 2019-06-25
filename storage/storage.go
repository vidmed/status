package storage

import (
	"context"

	"github.com/vidmed/status/dto"
)

type Storage interface {
	GetStatus(ctx context.Context, req *dto.OrderStatusRequest) (*dto.OrderStatus, error)
	GetStatuses(ctx context.Context, req *dto.OrderStatusesRequest) (*dto.OrderStatuses, error)
	GetStatusHistory(ctx context.Context, req *dto.StatusHistoryRequest) (*dto.StatusHistory, error)
}
