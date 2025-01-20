package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xan-mortum/infrastructure/resourcesapi/store"
)

type Users struct {
	Store *store.Users
}

func NewUsers(store *store.Users) *Users {
	return &Users{
		Store: store,
	}
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
func (b *Users) GetUsersHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, b.Store.GetUsers())
}
