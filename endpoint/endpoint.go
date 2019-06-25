package endpoint

import (
	"context"

	"github.com/vidmed/status/dto"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/vidmed/status/service"
)

// Endpoints collects all of the endpoints that compose a status service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	GetStatusEndpoint        endpoint.Endpoint
	GetStatusesEndpoint      endpoint.Endpoint
	GetStatusHistoryEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.Service, mdw map[string][]endpoint.Middleware) Endpoints {
	getStatusEndpoint := MakeGetStatusEndpoint(s)
	getStatusEndpoint = LoggingMiddleware(log.With(logger, "method", "Sum"))(getStatusEndpoint)

	getStatusesEndpoint := MakeGetStatusesEndpoint(s)
	getStatusesEndpoint = LoggingMiddleware(log.With(logger, "method", "Sum"))(getStatusesEndpoint)

	getStatusHistoryEndpoint := MakeGetStatusHistoryEndpoint(s)
	getStatusHistoryEndpoint = LoggingMiddleware(log.With(logger, "method", "Sum"))(getStatusHistoryEndpoint)

	return Endpoints{
		GetStatusEndpoint:        getStatusEndpoint,
		GetStatusesEndpoint:      getStatusesEndpoint,
		GetStatusHistoryEndpoint: getStatusHistoryEndpoint,
	}
}

// MakeGetStatusEndpoint constructs a GetStatus endpoint wrapping the service.
func MakeGetStatusEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*dto.OrderStatusRequest)
		return s.GetStatus(ctx, req)
	}
}

// MakeGetStatusesEndpoint constructs a GetStatuses endpoint wrapping the service.
func MakeGetStatusesEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*dto.OrderStatusesRequest)
		return s.GetStatuses(ctx, req)
	}
}

// MakeGetStatusHistoryEndpoint constructs a GetStatusHistory endpoint wrapping the service.
func MakeGetStatusHistoryEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*dto.StatusHistoryRequest)
		return s.GetStatusHistory(ctx, req)
	}
}
