package controllers

import (
	"assignment-2/models"
	"assignment-2/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

	var OrderRequestBody models.OrderCreateRequestBody
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
	var OrderRequestBody models.OrderCreateRequestBody
	var items *models.Item
	var errItems error
	if err := ctx.BindJSON(&OrderRequestBody); err != nil {
		fmt.Println("! nil")

		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	for _, v := range OrderRequestBody.Items {
		item, err := oc.itemService.FindItem(v.ItemCode.String())
		if err != nil {
			fmt.Println("! nil")
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
			return
		}
		if v.ItemCode != item.ItemCode {
			fmt.Println("! nil")
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "can't find item code"})
			return
		}
	}
	orders, err := oc.orderService.CreateOrder(OrderRequestBody.CustomerName, OrderRequestBody.OrderedAt, OrderRequestBody.Items)
	for _, v := range OrderRequestBody.Items {
		item, err := oc.itemService.FindItem(v.ItemCode.String())
		if err != nil {
			fmt.Println("! nil")
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
			return
		}
		if v.ItemCode == item.ItemCode {
			items, errItems = oc.itemService.CreateItem(item.Description, item.Quantity, orders.ID)
			if errItems != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
				return
			}

		} else {
			fmt.Println("! nil")
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "can't find item code"})
			return
		}
	}
	if err != nil {
		fmt.Println("! nil")

		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"items": items, "orders": orders}})

}
