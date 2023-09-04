package controllers

import (
	"assignment-2/models"
	"assignment-2/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var ItemControllerInit *ItemController

type OrderController struct {
	orderService services.OrderService
	itemService  services.ItemService
}

func NewOrderController(orderService services.OrderService, itemService services.ItemService) *OrderController {
	return &OrderController{orderService, itemService}
}

func (oc *OrderController) GetAllOrders(ctx *gin.Context) {
	orders, err := oc.orderService.GetAllOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"items": orders}})
}

func (oc *OrderController) DeleteOrder(ctx *gin.Context) {
	orderID, _ := strconv.Atoi(ctx.Param("id"))

	response, err := oc.orderService.DeleteOrder(orderID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"orders": response}})
}

func (oc *OrderController) UpdateOrder(ctx *gin.Context) {
	orderID, _ := strconv.Atoi(ctx.Param("id"))

	var OrderRequestBody models.OrderRequestBody
	if err := ctx.BindJSON(&OrderRequestBody); err != nil {
		fmt.Println("! nil")

		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	fmt.Println("OrderRequestBody", orderID, OrderRequestBody.CustomerName, OrderRequestBody.OrderedAt, OrderRequestBody.Items)

	response, err := oc.orderService.UpdateOrder(orderID, OrderRequestBody.CustomerName, OrderRequestBody.OrderedAt, OrderRequestBody.Items)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"orders": response}})
}

func (oc *OrderController) CreateOrder(ctx *gin.Context) {
	var OrderRequestBody models.OrderRequestBody
	if err := ctx.BindJSON(&OrderRequestBody); err != nil {
		fmt.Println("! nil")

		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	orders, err := oc.orderService.CreateOrder(OrderRequestBody.CustomerName, OrderRequestBody.OrderedAt, OrderRequestBody.Items)
	if err != nil {
		fmt.Println("! nil")

		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"items": orders}})

}
