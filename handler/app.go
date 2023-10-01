package handler

import (
	"assignment-2/docs"
	"assignment-2/infra/database"
	"assignment-2/repository/item_repository/item_pg"
	"assignment-2/repository/order_repository/order_pg"
	"assignment-2/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func StartApp() {
	database.InitiliazeDatabase()

	db := database.GetDatabaseInstance()

	orderRepo := order_pg.NewOrderPG(db)

	itemRepo := item_pg.NewItemPG(db)

	orderService := service.NewOrderService(orderRepo, itemRepo)

	orderHandler := NewOrderHandler(orderService)

	route := gin.Default()

	docs.SwaggerInfo.Title = "Assgnment-2"
	docs.SwaggerInfo.Description = "Tugas ke-2 dari Hacktiv8"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http"}
	
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	route.POST("/orders", orderHandler.CreateOrder)

	route.GET("/orders", orderHandler.ReadOrders)

	route.PUT("/orders/:orderId", orderHandler.UpdateOrder)

	route.DELETE("/orders/:orderId", orderHandler.DeleteOrder)

	route.Run(":8080")
}
