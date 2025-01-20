package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"xan-mortum/infrastructure/authapi/cmd/docs"
	"xan-mortum/infrastructure/authapi/config"
	"xan-mortum/infrastructure/authapi/internal/handler"
)

func main() {
	cfg := config.GetConfig()
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	tokenHandler := handler.NewTokenHandler(cfg.JWT)

	v1 := r.Group("/api/v1")
	v1.POST("/token", tokenHandler.GenerateToken)
	v1.POST("/check_token", tokenHandler.CheckToken)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(cfg.Server.Port)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
