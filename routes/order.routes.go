package routes

import (
	"assignment-2/controllers"
	"assignment-2/services"

	"github.com/gin-gonic/gin"
)

type OrderRouteController struct {
	orderController *controllers.OrderController
}

func NewOrderRouteController(orderController *controllers.OrderController) *OrderRouteController {
	return &OrderRouteController{orderController}
}

func (oc *OrderRouteController) OrderController(rg *gin.RouterGroup, orderService services.OrderService) {
	router := rg.Group("orders")
	router.GET("", oc.orderController.GetAllOrders)
	router.POST("", oc.orderController.CreateOrder)
}
