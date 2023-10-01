package service

import (
	"assignment-2/dto"
	"assignment-2/entity"
	"assignment-2/pkg/errs"
	"assignment-2/repository/item_repository"
	"assignment-2/repository/order_repository"
	"fmt"
	"net/http"
)

type orderService struct {
	OrderRepo order_repository.Repository
	ItemRepo  item_repository.Repository
}

type OrderService interface {
	CreateOrder(newOrderRequest dto.NewOrderRequest) (*dto.NewOrderResponse, errs.Error)
	ReadOrders() (*dto.GetOrdersResponse, errs.Error)
	UpdateOrder(orderId int, newOrderRequest dto.NewOrderRequest) (*dto.NewOrderResponse, errs.Error)
	DeleteOrder(orderId int) (*dto.DeleteOrdersResponse, errs.Error)
}

func NewOrderService(orderRepo order_repository.Repository, itemRepo item_repository.Repository) OrderService {
	return &orderService{
		OrderRepo: orderRepo,
		ItemRepo:  itemRepo,
	}
}

func (os *orderService) ReadOrders() (*dto.GetOrdersResponse, errs.Error) {
	orders, err := os.OrderRepo.ReadOrders()

	if err != nil {
		return nil, err
	}

	orderResult := []dto.OrderWithItems{}

	for _, eachOrder := range orders {
		order := dto.OrderWithItems{
			OrderId:      eachOrder.Order.OrderId,
			CustomerName: eachOrder.Order.CustomerName,
			OrderedAt:    eachOrder.Order.OrderedAt,
			CreatedAt:    eachOrder.Order.CreatedAt,
			UpdatedAt:    eachOrder.Order.UpdatedAt,
			Items:        []dto.GetItemResponse{},
		}

		for _, eachItem := range eachOrder.Items {
			item := dto.GetItemResponse{
				ItemId:      eachItem.ItemId,
				ItemCode:    eachItem.ItemCode,
				Quantity:    eachItem.Quantity,
				Description: eachItem.Description,
				OrderId:     eachItem.OrderId,
				CreatedAt:   eachItem.CreatedAt,
				UpdatedAt:   eachItem.UpdatedAt,
			}
			order.Items = append(order.Items, item)
		}
		orderResult = append(orderResult, order)
	}

	response := dto.GetOrdersResponse{
		StatusCode: http.StatusOK,
		Message:    "Order succesfully fetched",
		Data:       orderResult,
	}
	return &response, nil
}

func (os *orderService) DeleteOrder(orderId int) (*dto.DeleteOrdersResponse, errs.Error) {
	// Mengecek apakah orderId yang diberikan client itu ada jika tidak berikan error
	_, err := os.OrderRepo.ReadOrderById(orderId)

	if err != nil {
		return nil, err
	}
	// jika ada hapus order menggunakan orderId
	err = os.OrderRepo.DeleteOrder(orderId)

	if err != nil {
		return nil, err
	}

	response := dto.DeleteOrdersResponse{
		StatusCode: http.StatusOK,
		Message:    "Order succesfully Deleted",
	}
	return &response, nil
}

func (os *orderService) UpdateOrder(orderId int, newOrderRequest dto.NewOrderRequest) (*dto.NewOrderResponse, errs.Error) {

	_, err := os.OrderRepo.ReadOrderById(orderId)

	if err != nil {
		return nil, err
	}

	itemCodes := []any{}

	for _, eachItem := range newOrderRequest.Items {
		itemCodes = append(itemCodes, eachItem.ItemCode)
	}

	items, err := os.ItemRepo.GetItemsByCodes(itemCodes)

	if err != nil {
		return nil, err
	}

	for _, eachItemFromRequest := range newOrderRequest.Items {
		isFound := false

		for _, eachItem := range items {

			if eachItem.OrderId != orderId {

				return nil, errs.NewBadRequest(fmt.Sprintf("item with item code %s doesn't belong to the order with id %d", eachItem.ItemCode, orderId))

			}

			if eachItemFromRequest.ItemCode == eachItem.ItemCode {
				isFound = true
				break
			}
		}

		if !isFound {
			return nil, errs.NewNotFoundError(fmt.Sprintf("item with item code %s is not found", eachItemFromRequest.ItemCode))
		}
	}

	itemPayload := []entity.Item{}

	for _, eachItemFromRequest := range newOrderRequest.Items {
		item := entity.Item{
			ItemCode:    eachItemFromRequest.ItemCode,
			Description: eachItemFromRequest.Description,
			Quantity:    eachItemFromRequest.Quantity,
		}

		itemPayload = append(itemPayload, item)

	}

	orderPayload := entity.Order{
		OrderId:      orderId,
		OrderedAt:    newOrderRequest.OrderedAt,
		CustomerName: newOrderRequest.CustomerName,
	}

	err = os.OrderRepo.UpdateOrder(orderPayload, itemPayload)

	if err != nil {
		return nil, err
	}

	response := dto.NewOrderResponse{
		StatusCode: http.StatusOK,
		Message:    "order successfully updated",
		Data:       nil,
	}

	return &response, nil
}

func (os *orderService) CreateOrder(newOrderRequest dto.NewOrderRequest) (*dto.NewOrderResponse, errs.Error) {
	// Request
	orderPayLoad := entity.Order{
		OrderedAt:    newOrderRequest.OrderedAt,
		CustomerName: newOrderRequest.CustomerName,
	}

	itemPayLoad := []entity.Item{}

	for _, eachItem := range newOrderRequest.Items {
		item := entity.Item{
			ItemCode:    eachItem.ItemCode,
			Description: eachItem.Description,
			Quantity:    eachItem.Quantity,
		}
		itemPayLoad = append(itemPayLoad, item)
	}

	err := os.OrderRepo.CreateOrder(orderPayLoad, itemPayLoad)

	if err != nil {
		return nil, err
	}
	// Response
	response := dto.NewOrderResponse{
		StatusCode: http.StatusCreated,
		Message:    "new order successfully created",
		Data:       nil,
	}

	return &response, nil
}
