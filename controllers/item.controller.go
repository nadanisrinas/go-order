package controllers

import (
	"assignment-2/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemController struct {
	itemService services.ItemService
}

func NewItemController(itemService services.ItemService) *ItemController {
	return &ItemController{itemService}
}

func (ic *ItemController) GetAllItems(ctx *gin.Context) {
	items, err := ic.itemService.GetAllItems()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"items": items}})
}

func (ic *ItemController) CreateItem(ctx *gin.Context) {
	// var item models.Item
	itemsResponse, err := ic.itemService.CreateItem("desc1", 3)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"items": itemsResponse}})
}
