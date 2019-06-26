package endpoint

import (
	endpointDTO "github.com/vidmed/status/pkg/endpoint/dto"
	servicetDTO "github.com/vidmed/status/pkg/service/dto"
)

func convertOrderStatusRequest(req *endpointDTO.OrderStatusRequest) *servicetDTO.OrderStatusRequest {
	return &servicetDTO.OrderStatusRequest{
		OrderId: req.OrderId,
	}
}

func convertOrderStatusesRequest(req *endpointDTO.OrderStatusesRequest) *servicetDTO.OrderStatusesRequest {
	return &servicetDTO.OrderStatusesRequest{
		OrderIds: req.OrderIds,
	}
}

func convertStatusHistoryRequest(req *endpointDTO.StatusHistoryRequest) *servicetDTO.StatusHistoryRequest {
	return &servicetDTO.StatusHistoryRequest{
		OrderId: req.OrderId,
		Limit:   req.Limit,
		Offset:  req.Offset,
		Filter:  req.Filter,
	}
}

func convertOrderStatusResponse(req *servicetDTO.OrderStatus) *endpointDTO.OrderStatusResponse {
	return &endpointDTO.OrderStatusResponse{
		OrderInfo: convertOrderInfo(req.OrderInfo),
		Status:    convertStatus(req.Status),
	}
}

func convertOrderStatusesResponse(req *servicetDTO.OrderStatuses) *endpointDTO.OrderStatusesResponse {
	return &endpointDTO.OrderStatusesResponse{
		SucceedOrders: convertSucceedOrders(req.SucceedOrders),
		FailedOrders:  convertFailedOrders(req.FailedOrders),
	}
}

func convertStatusHistoryResponse(req *servicetDTO.StatusHistory) *endpointDTO.StatusHistoryResponse {
	return &endpointDTO.StatusHistoryResponse{
		OrderInfo: convertOrderInfo(req.OrderInfo),
		Rows:      convertStatuses(req.Rows),
		Meta:      convertMeta(req.Meta),
	}
}

func convertOrderInfo(oi *servicetDTO.OrderInfo) *endpointDTO.OrderInfo {
	return &endpointDTO.OrderInfo{
		OrderId:        oi.OrderId,
		ProviderNumber: oi.ProviderNumber,
		Barcode:        oi.Barcode,
		ClientNumber:   oi.ClientNumber,
	}
}

func convertStatus(s *servicetDTO.Status) *endpointDTO.Status {
	return &endpointDTO.Status{
		Key:                 s.Key,
		Name:                s.Name,
		Description:         s.Description,
		Created:             s.Created,
		ProviderCode:        s.ProviderCode,
		ProviderName:        s.ProviderName,
		ProviderDescription: s.ProviderDescription,
		CreatedProvider:     s.CreatedProvider,
		ErrorCode:           s.ErrorCode,
	}
}

func convertSucceedOrders(orders []*servicetDTO.OrderStatus) []*endpointDTO.OrderStatusResponse {
	res := make([]*endpointDTO.OrderStatusResponse, len(orders))
	for i, order := range orders {
		res[i] = convertOrderStatusResponse(order)
	}

	return res
}

func convertFailedOrder(order *servicetDTO.FailedOrder) *endpointDTO.FailedOrder {
	return &endpointDTO.FailedOrder{
		OrderId: order.OrderId,
		Message: order.Message,
	}
}

func convertFailedOrders(orders []*servicetDTO.FailedOrder) []*endpointDTO.FailedOrder {
	res := make([]*endpointDTO.FailedOrder, len(orders))
	for i, order := range orders {
		res[i] = convertFailedOrder(order)
	}

	return res
}

func convertStatuses(statuses []*servicetDTO.Status) []*endpointDTO.Status {
	res := make([]*endpointDTO.Status, len(statuses))
	for i, status := range statuses {
		res[i] = convertStatus(status)
	}

	return res
}

func convertMeta(meta *servicetDTO.Meta) *endpointDTO.Meta {
	return &endpointDTO.Meta{
		Total:  meta.Total,
		Offset: meta.Offset,
		Limit:  meta.Limit,
	}
}
