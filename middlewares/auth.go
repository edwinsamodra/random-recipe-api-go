package middlewares

import (
	"fmt"
	"net/http"
	"webapp1/util"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	const API_KEY = "BaksoBeranak"
	const HEADER_KEY = "api-key-secret"

	return func(c *gin.Context) {
		apiKeySecret := c.GetHeader(HEADER_KEY)

		if apiKeySecret != API_KEY {
			fmt.Println(apiKeySecret)
			c.AbortWithStatusJSON(http.StatusUnauthorized, util.BuildResponse("error", nil))
			return
		}

		c.Next()
	}
}
