package controllers

import (
	"assignment-2/models"
	"assignment-2/services"
	"log"
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
	var itemRequestBody models.ItemRequestBody
	if err := ctx.BindJSON(&itemRequestBody); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	// var item models.Item
	itemsResponse, err := ic.itemService.CreateItem(itemRequestBody.Description, itemRequestBody.Quantity)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"items": itemsResponse}})
}

func (ic *ItemController) FindItem(ItemCode string) *models.Item {
	// var item models.Item
	itemsResponse, err := ic.itemService.FindItem(ItemCode)
	if err != nil {
		log.Fatal("err find item", err)
	}

	return &itemsResponse
}
func FindItem(ItemCode string) *models.Item {
	// var item models.Item
	itemsResponse, err := services.FindItem(ItemCode)
	if err != nil {
		log.Fatal("err find item", err)
	}

	return &itemsResponse
}
