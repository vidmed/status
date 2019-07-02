package dummy

import (
	"context"
	"strconv"

	"github.com/vidmed/status/pkg/service/dto"

	"github.com/vidmed/status/pkg/service/storage"
)

//var _ = storage.Storage(dummyStorage{})

type dummyStorage struct {
}

func NewDummyStorage() storage.Storage {
	return dummyStorage{}
}

func (ds dummyStorage) Close() {

}

func (ds dummyStorage) GetStatus(ctx context.Context, req *dto.OrderStatusRequest) (*dto.OrderStatus, error) {
	return &dto.OrderStatus{
		OrderInfo: ds.orderInfo(req.OrderId),
		Status:    ds.status(),
	}, nil
}

func (ds dummyStorage) GetStatuses(ctx context.Context, req *dto.OrderStatusesRequest) (*dto.OrderStatuses, error) {
	orderStatuses := &dto.OrderStatuses{}

	for _, orderID := range req.OrderIds {
		if orderStatus, err := ds.GetStatus(ctx, &dto.OrderStatusRequest{OrderId: orderID}); err != nil {
			orderStatuses.SucceedOrders = append(orderStatuses.SucceedOrders, orderStatus)
		} else {
			orderStatuses.FailedOrders = append(orderStatuses.FailedOrders, ds.failedOrder(orderID))
		}

	}

	return orderStatuses, nil
}

func (ds dummyStorage) GetStatusHistory(ctx context.Context, req *dto.StatusHistoryRequest) (*dto.StatusHistory, error) {
	return &dto.StatusHistory{
		OrderInfo: ds.orderInfo(req.OrderId),
		Rows:      []*dto.Status{ds.status()},
		Meta:      &dto.Meta{Total: 1, Offset: req.Offset, Limit: req.Limit},
	}, nil
}

func (ds dummyStorage) orderInfo(orderID int) *dto.OrderInfo {
	return &dto.OrderInfo{
		OrderId:        strconv.Itoa(orderID),
		ProviderNumber: "string",
		Barcode:        "string",
		ClientNumber:   "string",
	}
}

func (ds dummyStorage) status() *dto.Status {
	return &dto.Status{
		Key:                 "pending",
		Name:                "Не обработан",
		Description:         "Заказ ещё не обработан службой доставки",
		Created:             "2015-08-07T10:15:17+04:00",
		ProviderCode:        "4",
		ProviderName:        "string",
		ProviderDescription: "string",
		CreatedProvider:     "2015-08-07T10:15:25+04:00",
		ErrorCode:           "200",
	}
}

func (ds dummyStorage) failedOrder(orderID int) *dto.FailedOrder {
	return &dto.FailedOrder{
		OrderId: orderID,
		Message: "error message",
	}
}
