package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"xan-mortum/infrastructure/resourcesapi/cmd/docs"
	"xan-mortum/infrastructure/resourcesapi/config"
	"xan-mortum/infrastructure/resourcesapi/internal/handler"
	"xan-mortum/infrastructure/resourcesapi/store"
)

func main() {
	cfg := config.GetConfig()

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	users := handler.NewUsers(store.NewUsers())
	books := handler.NewBooks(store.NewBooks())

	v1 := r.Group("/api/v1")
	v1.GET("/users", users.GetUsersHandler)
	v1.GET("/books", books.GetBooksHandler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(cfg.Server.Port)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
