package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/vidmed/status/pkg/service/dto"
)

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

func (mw loggingMiddleware) GetStatus(ctx context.Context, req *dto.OrderStatusRequest) (res *dto.OrderStatus, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "GetStatus",
			"input", req,
			"output", res,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	res, err = mw.next.GetStatus(ctx, req)
	return
}

func (mw loggingMiddleware) GetStatuses(ctx context.Context, req *dto.OrderStatusesRequest) (res *dto.OrderStatuses, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "GetStatuses",
			"input", req,
			"output", res,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	res, err = mw.next.GetStatuses(ctx, req)
	return
}

func (mw loggingMiddleware) GetStatusHistory(ctx context.Context, req *dto.StatusHistoryRequest) (res *dto.StatusHistory, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "GetStatusHistory",
			"input", req,
			"output", res,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	res, err = mw.next.GetStatusHistory(ctx, req)
	return
}
