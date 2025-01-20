package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xan-mortum/infrastructure/apigateway/internal/service"
)

type Auth struct {
	Client service.AuthClient
}

func NewAuth(client service.AuthClient) *Auth {
	return &Auth{
		Client: client,
	}
}

// @BasePath /api/v1

// SignInHandler	godoc
//
//	@Summary		Sing In
//	@Description	SignIn
//	@Produce		application/json
//	@Tags			auth
//	@Success		200	{object} string
//	@Router			/sign_in [post]
func (a *Auth) SignInHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	res, err := a.Client.SignIn()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (a *Auth) AuthorisationMiddleware(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	token := c.GetHeader("Authorization")
	res, err := a.Client.CheckToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if "true" != string(res) {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}

	c.Next()
}
