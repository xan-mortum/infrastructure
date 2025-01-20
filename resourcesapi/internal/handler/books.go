package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xan-mortum/infrastructure/resourcesapi/store"
)

type Books struct {
	Store *store.Books
}

func NewBooks(store *store.Books) *Books {
	return &Books{
		Store: store,
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
func (b *Books) GetBooksHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, b.Store.GetBooks())
}
