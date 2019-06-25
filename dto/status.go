package dto

// ------------------ requests

type OrderStatusRequest struct {
	OrderId int `json:"orderId"`
}

type OrderStatusesRequest struct {
	OrderIds []int `json:"orderIds"`
}

type StatusHistoryRequest struct {
	OrderId int    `json:"orderId"`
	Limit   int    `json:"limit"`
	Offset  int    `json:"offset"`
	Filter  string `json:"filter"`
}

// ------------------ responses

type OrderStatus struct {
	OrderInfo *OrderInfo `json:"orderInfo"`
	Status    *Status    `json:"status"`
}

type OrderStatuses struct {
	SucceedOrders []*OrderStatus `json:"succeedOrders"`
	FailedOrders  []*FailedOrder `json:"failedOrders"`
}

type StatusHistory struct {
	OrderInfo *OrderInfo `json:"orderInfo"`
	Rows      []*Status  `json:"rows"`
	Meta      *Meta      `json:"meta"`
}

type Status struct {
	Key                 string `json:"key"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	Created             string `json:"created"`
	ProviderCode        string `json:"providerCode"`
	ProviderName        string `json:"providerName"`
	ProviderDescription string `json:"providerDescription"`
	CreatedProvider     string `json:"createdProvider"`
	ErrorCode           string `json:"errorCode"`
}

type OrderInfo struct {
	OrderId        string `json:"orderId"`
	ProviderNumber string `json:"providerNumber"`
	Barcode        string `json:"barcode"`
	ClientNumber   string `json:"clientNumber"`
}

type FailedOrder struct {
	OrderId int    `json:"orderId"`
	Message string `json:"message"`
}

type Meta struct {
	Total  int `json:"total"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
