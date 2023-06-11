package middlewares

import (
	"net/http"
	"strings"
	"webapp1/util"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

var jwtKey = []byte("secret-key")

// func Auth() gin.HandlerFunc {
// 	const API_KEY = "BaksoBeranak"
// 	const HEADER_KEY = "api-key-secret"

// 	return func(c *gin.Context) {
// 		apiKeySecret := c.GetHeader(HEADER_KEY)

// 		if apiKeySecret != API_KEY {
// 			fmt.Println(apiKeySecret)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, util.BuildResponse("error", nil))
// 			return
// 		}

// 		c.Next()
// 	}
// }

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, util.BuildResponse("error", nil))
			return
		}

		token := strings.Split(authorization, " ")

		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(token[1], claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, util.BuildResponse("error", nil))
				return
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, util.BuildResponse("error", nil))
			return
		}

		if !tkn.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, util.BuildResponse("error", nil))
			return
		}

		c.Next()
	}
}
