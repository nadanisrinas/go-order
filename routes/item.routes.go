package routes

import (
	"assignment-2/controllers"
	"assignment-2/services"

	"github.com/gin-gonic/gin"
)

type ItemRouteController struct {
	itemController *controllers.ItemController
}

func NewItemRouteController(itemController *controllers.ItemController) *ItemRouteController {
	return &ItemRouteController{itemController}
}

func (ic *ItemRouteController) ItemRoute(rg *gin.RouterGroup, itemService services.ItemService) {
	router := rg.Group("items")
	router.GET("", ic.itemController.GetAllItems)
	router.POST("", ic.itemController.CreateItem)
}
