package endpoint

import (
	"context"

	endpointDTO "github.com/vidmed/status/pkg/endpoint/dto"
	"github.com/vidmed/status/pkg/service"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
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
func New(s service.Service, logger log.Logger) Endpoints {
	getStatusEndpoint := MakeGetStatusEndpoint(s)
	getStatusEndpoint = LoggingMiddleware(log.With(logger, "method", "GetStatus"))(getStatusEndpoint)

	getStatusesEndpoint := MakeGetStatusesEndpoint(s)
	getStatusesEndpoint = LoggingMiddleware(log.With(logger, "method", "GetStatuses"))(getStatusesEndpoint)

	getStatusHistoryEndpoint := MakeGetStatusHistoryEndpoint(s)
	getStatusHistoryEndpoint = LoggingMiddleware(log.With(logger, "method", "GetStatusHistory"))(getStatusHistoryEndpoint)

	return Endpoints{
		GetStatusEndpoint:        getStatusEndpoint,
		GetStatusesEndpoint:      getStatusesEndpoint,
		GetStatusHistoryEndpoint: getStatusHistoryEndpoint,
	}
}

// MakeGetStatusEndpoint constructs a GetStatus endpoint wrapping the service.
func MakeGetStatusEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*endpointDTO.OrderStatusRequest)
		serviceRequest := convertOrderStatusRequest(req)

		resp, err := s.GetStatus(ctx, serviceRequest)
		if err != nil {
			return nil, err
		}

		return convertOrderStatusResponse(resp), nil
	}
}

// MakeGetStatusesEndpoint constructs a GetStatuses endpoint wrapping the service.
func MakeGetStatusesEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*endpointDTO.OrderStatusesRequest)
		serviceRequest := convertOrderStatusesRequest(req)

		resp, err := s.GetStatuses(ctx, serviceRequest)
		if err != nil {
			return nil, err
		}

		return convertOrderStatusesResponse(resp), nil
	}
}

// MakeGetStatusHistoryEndpoint constructs a GetStatusHistory endpoint wrapping the service.
func MakeGetStatusHistoryEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*endpointDTO.StatusHistoryRequest)
		serviceRequest := convertStatusHistoryRequest(req)

		resp, err := s.GetStatusHistory(ctx, serviceRequest)
		if err != nil {
			return nil, err
		}

		return convertStatusHistoryResponse(resp), nil
	}
}
