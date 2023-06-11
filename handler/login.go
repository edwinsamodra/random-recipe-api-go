package handler

import (
	"net/http"
	"time"
	"webapp1/model"
	"webapp1/util"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func HandlerLogin(c *gin.Context) {
	var creds model.Credentials
	err := c.BindJSON(&creds)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, util.BuildResponse("error", nil))
		return
	}

	expectedPassword, ok := util.Users[creds.Username]
	if !ok || expectedPassword.Password != creds.Password {
		c.AbortWithStatusJSON(http.StatusBadRequest, util.BuildResponse("error", nil))
		return
	}

	exp := time.Now().Add(5 * time.Minute)
	claims := &util.Claims{
		Username: creds.Username,
		Role:     expectedPassword.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(util.JwtKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, util.BuildResponse("error", nil))
		return
	}

	dataToken := make(map[string]any)
	dataToken["token"] = tokenString
	dataToken["expires"] = exp

	c.JSON(http.StatusOK, util.BuildResponse("success logged in", dataToken))
}
