package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"xan-mortum/infrastructure/authapi/config"
	"xan-mortum/infrastructure/authapi/internal/command"
)

type TokenHandler struct {
	config *config.JWT
}

func NewTokenHandler(config *config.JWT) *TokenHandler {
	return &TokenHandler{
		config: config,
	}
}

// @BasePath /api/v1

// GenerateToken	godoc
//
//	@Summary		Generate new jwt token
//	@Description	Generate new jwt token
//	@Produce		application/json
//	@Tags			token
//	@Success		200	{object} string
//	@Router			/token [post]
func (t *TokenHandler) GenerateToken(c *gin.Context) {
	key := []byte(t.config.SecretKey)
	jwtToken := jwt.New(jwt.SigningMethodHS256)
	c.Header("Content-Type", "application/json")
	token, err := jwtToken.SignedString(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, token)
}

// @BasePath /api/v1

// CheckToken		godoc
//
//	@Summary		Check token
//	@Description	Check token
//	@Param  		token body command.CheckToken true "Some token"
//	@Produce		application/json
//	@Tags			token
//	@Success		200	{object} bool
//	@Router			/check_token [post]
func (t *TokenHandler) CheckToken(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	cmd := &command.CheckToken{}
	err := c.BindJSON(cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if cmd.Token == "" {
		c.JSON(http.StatusOK, true)
		return
	}

	c.JSON(http.StatusOK, true)
}
