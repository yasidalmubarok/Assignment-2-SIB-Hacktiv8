package handler

import (
	"assignment-2/dto"
	"assignment-2/pkg/errs"
	"assignment-2/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	OrderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) orderHandler {
	return orderHandler{
		OrderService: orderService,
	}
}

// GetOrders godoc
// @Summary ping example
// @Schemes
// @Description Get Order with Item Data
// @ID get-order-datas
// @Tags orders
// @Produce json
// @Success 200 {object} dto.GetOrdersResponse
// @Router /orders [get]
func (oh *orderHandler) ReadOrders(ctx *gin.Context) {
	response, err := oh.OrderService.ReadOrders()

	if err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid json request body")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// UpdateOrders godoc
// @Summary ping example
// @Schemes
// @Description Update Order Data by Id
// @ID update-order-with-item-data
// @Tags orders
// @Accept json
// @Produce json
// @Param RequestBody body dto.NewOrderRequest true "request body json"
// @Param orderId path int true "order's id"
// @Success 200 {object} dto.NewOrderResponse
// @Router /orders/{orderId} [put]
func (oh *orderHandler) UpdateOrder(ctx *gin.Context) {
	var newOrderRequest dto.NewOrderRequest

	var orderId, _ = strconv.Atoi(ctx.Param("orderId"))

	if err := ctx.ShouldBindJSON(&newOrderRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid json request body")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := oh.OrderService.UpdateOrder(orderId, newOrderRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)

		return
	}

	ctx.JSON(response.StatusCode, response)
}

// DeleteOrders godoc
// @Summary ping example
// @Schemes
// @Description Delete Order Data by Id
// @ID delete-order-data
// @Tags orders
// @Produce json
// @Param orderId path int true "orders id"
// @Success 200 {object} dto.DeleteOrdersResponse
// @Router /orders/{orderId} [delete]
func (oh *orderHandler) DeleteOrder(ctx *gin.Context) {
	orderId, err := strconv.Atoi(ctx.Param("orderId"))

	if err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid json request body")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := oh.OrderService.DeleteOrder(orderId)

	if err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid json request body")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// CreateOrder godoc
// @Summary ping example
// @Schemes
// @Description Create Order Data
// @ID create-new-order
// @Tags orders
// @Accept json
// @Produce json
// @Param RequestBody body dto.NewOrderRequest true "request body json"
// @Success 201 {object} dto.NewOrderResponse
// @Router /orders [post]
func (oh *orderHandler) CreateOrder(ctx *gin.Context) {
	var newOrderRequest dto.NewOrderRequest

	if err := ctx.ShouldBindJSON(&newOrderRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid json request body")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := oh.OrderService.CreateOrder(newOrderRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(response.StatusCode, response)
}
