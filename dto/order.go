package dto

import "time"

//Request
type NewOrderRequest struct {
	OrderedAt    time.Time        `json:"orderedAt" example:"2023-07-10T21:21:46+00:00"`
	CustomerName string           `json:"customerName"`
	Items        []NewItemRequest `json:"items"`
}

//Response
type OrderWithItems struct {
	OrderId      int               `json:"orderId"`
	CustomerName string            `json:"customerName"`
	OrderedAt    time.Time         `json:"orderedAt"`
	CreatedAt    time.Time         `json:"createdAt"`
	UpdatedAt    time.Time         `json:"updatedAt"`
	Items        []GetItemResponse `json:"items"`
}

type GetOrdersResponse struct {
	StatusCode int              `json:"statusCode"`
	Message    string           `json:"message"`
	Data       []OrderWithItems `json:"data"`
}

type NewOrderResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type DeleteOrdersResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
