package controllers

import (
	"assignment-2/models"
	"assignment-2/services"
	"fmt"
	"net/http"

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

func (oc *OrderController) CreateOrder(ctx *gin.Context) {
	var OrderRequestBody models.OrderRequestBody
	if err := ctx.BindJSON(&OrderRequestBody); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	fmt.Println("ooo", OrderRequestBody.Items)
	for i, v := range OrderRequestBody.Items {
		fmt.Println("=======>", v.ItemCode)
		fmt.Println("=======>", &v.ItemCode)

		findOrderResponse := FindItem((v.ItemCode).String())
		fmt.Println("findOrderResponse", findOrderResponse)
		if findOrderResponse != nil {
			_, err := (*&oc.itemService).CreateItem(OrderRequestBody.Items[i].Description, OrderRequestBody.Items[i].Quantity)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})

			}
			orders, err := oc.orderService.CreateOrder(OrderRequestBody.OrderedAt, OrderRequestBody.Items)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})

			}
			ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"orders": orders}})
		} else {
			fmt.Println("pppp")
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "error say"})

		}
	}

}
