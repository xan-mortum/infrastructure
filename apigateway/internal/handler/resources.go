package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xan-mortum/infrastructure/apigateway/internal/service"
)

type Resources struct {
	Client service.ResourceClient
}

func NewResources(client service.ResourceClient) *Resources {
	return &Resources{
		Client: client,
	}
}

// @BasePath /api/v1

// GetBooksHandler godoc
// @Summary get books
// @Schemes
// @Description get books
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {string} map[int]string
// @Router /books [get]
func (b *Resources) GetBooksHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	res, err := b.Client.GetBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": string(res),
	})
}

// @BasePath /api/v1

// GetUsersHandler godoc
// @Summary get users
// @Schemes
// @Description get users
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {string} map[int]string
// @Router /users [get]
func (b *Resources) GetUsersHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	res, err := b.Client.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": string(res),
	})
}
