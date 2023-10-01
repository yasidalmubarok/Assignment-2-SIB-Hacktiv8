package dto

import "time"
//Request
type NewItemRequest struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
//Resposnse
type GetItemResponse struct {
	ItemId      int       `json:"itemId"`
	ItemCode    string    `json:"itemCode"`
	Quantity    int       `json:"quantity"`
	Description string    `json:"description"`
	OrderId     int       `json:"orderId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
