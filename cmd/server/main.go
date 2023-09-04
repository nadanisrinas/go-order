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
	server               *gin.Engine
	host                 = "localhost"
	port                 = 5432
	user                 = "postgres"
	password             = "root"
	dbname               = "assignment2"
	db                   *gorm.DB
	itemService          services.ItemService
	orderService         services.OrderService
	itemRouteController  *routes.ItemRouteController
	orderRouteController *routes.OrderRouteController
	itemController       *controllers.ItemController
	orderController      *controllers.OrderController
	err                  error
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("err conn to db: ", err)

	}
	errMigrate := db.Debug().AutoMigrate(models.Order{}, models.Item{})
	if errMigrate != nil {
		log.Fatal("error migrate db", errMigrate)
	}
	fmt.Println("db successfully migrated...")
}

func init() {
	StartDB()
	//init service
	itemService = services.NewItemService(db)
	orderService = services.NewOrderService(db)
	//init controller
	itemController = controllers.NewItemController(itemService)
	orderController = controllers.NewOrderController(orderService, itemService)
	//initroutes
	itemRouteController = routes.NewItemRouteController(itemController)
	orderRouteController = routes.NewOrderRouteController(orderController)
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
	orderRouteController.OrderRoute(router, orderService)
	fmt.Println("routes running")
	server.Run(":" + "8080")

}
