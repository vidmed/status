package dto

// ------------------ requests

type OrderStatusRequest struct {
	OrderId int
}

type OrderStatusesRequest struct {
	OrderIds []int
}

type StatusHistoryRequest struct {
	OrderId int
	Limit   int
	Offset  int
	Filter  string
}

// ------------------ responses

type OrderStatus struct {
	OrderInfo *OrderInfo
	Status    *Status
}

type OrderStatuses struct {
	SucceedOrders []*OrderStatus
	FailedOrders  []*FailedOrder
}

type StatusHistory struct {
	OrderInfo *OrderInfo
	Rows      []*Status
	Meta      *Meta
}

type Status struct {
	Key                 string
	Name                string
	Description         string
	Created             string
	ProviderCode        string
	ProviderName        string
	ProviderDescription string
	CreatedProvider     string
	ErrorCode           string
}

type OrderInfo struct {
	OrderId        string
	ProviderNumber string
	Barcode        string
	ClientNumber   string
}

type FailedOrder struct {
	OrderId int
	Message string
}

type Meta struct {
	Total  int
	Offset int
	Limit  int
}
