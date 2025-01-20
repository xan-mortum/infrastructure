package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"xan-mortum/infrastructure/apigateway/cmd/docs"
	"xan-mortum/infrastructure/apigateway/config"
	"xan-mortum/infrastructure/apigateway/internal/handler"
	"xan-mortum/infrastructure/apigateway/transport"
)

func main() {
	cfg := config.GetConfig()

	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	authHandler := handler.NewAuth(transport.NewAuthHttpClient(cfg.Auth))
	resourcesHandler := handler.NewResources(transport.NewResourcesHttpClient(cfg.Resources))

	v1 := r.Group("/api/v1")
	v1.POST("/sign_in", authHandler.SignInHandler)
	v1.GET("/users", authHandler.AuthorisationMiddleware, resourcesHandler.GetUsersHandler)
	v1.GET("/books", resourcesHandler.GetBooksHandler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(cfg.Server.Port)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
