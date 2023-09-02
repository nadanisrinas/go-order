package main

import (
	"assignment-2/controllers"
	"assignment-2/models"
	"assignment-2/routes"
	"assignment-2/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	server              *gin.Engine
	host                = "localhost"
	port                = 5432
	user                = "postgres"
	password            = "root"
	dbname              = "assignment2"
	db                  *gorm.DB
	itemService         services.ItemService
	itemRouteController *routes.ItemRouteController
	itemController      *controllers.ItemController
	err                 error
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("err conn to db: ", err)

	}
	db.Debug().AutoMigrate(models.Item{}, models.Order{})
	fmt.Println("db successfully migrated...")
}

func init() {
	StartDB()
	//init service
	itemService = services.NewItemService(db)
	//init controller
	itemController = controllers.NewItemController(itemService)
	//initroutes
	itemRouteController = routes.NewItemRouteController(itemController)
}

func GetDB() *gorm.DB {
	return db
}

func main() {

	server = gin.Default()
	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "helath check"})
	})
	itemRouteController.ItemRoute(router, itemService)
	fmt.Println("routes running")
	server.Run(":" + "8080")

}
