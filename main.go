package main

import (
	"net/http"
	"time"
	"webapp1/handler"
	"webapp1/middlewares"
	"webapp1/util"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// buat service yg mengembalikan resep (nama dan bahan2) dan harga total yg dibutuhkan:
// 1. ambil data dari https://api.spoonacular.com/recipes/random
// 2. bikin server
// 3. bikin endpoint dan handler nya
// 4. buat middleware check API Key

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"data":    nil,
		})
	})

	r.POST("/login", HandlerLogin)
	r.Use(middlewares.Auth())

	r.GET("/recipes", handler.HandlerGetRecipe)

	r.Run(":8080")
}

// start JWT
var jwtKey = []byte("secret-key")
var users = map[string]*User{
	"aditira": {
		Password: "password1",
		Role:     "admin",
	},
	"dito": {
		Password: "password2",
		Role:     "student",
	},
}

type User struct {
	Password string
	Role     string
}
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func HandlerLogin(c *gin.Context) {
	var creds Credentials
	err := c.BindJSON(&creds)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, util.BuildResponse("error", nil))
		return
	}

	expectedPassword, ok := users[creds.Username]
	if !ok || expectedPassword.Password != creds.Password {
		c.AbortWithStatusJSON(http.StatusBadRequest, util.BuildResponse("error", nil))
		return
	}

	exp := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		Role:     expectedPassword.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, util.BuildResponse("error", nil))
		return
	}

	dataToken := make(map[string]any)
	dataToken["token"] = tokenString
	dataToken["expires"] = exp

	c.JSON(http.StatusOK, util.BuildResponse("success logged in", dataToken))
}
