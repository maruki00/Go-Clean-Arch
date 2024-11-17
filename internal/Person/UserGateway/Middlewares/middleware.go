package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func middlewares() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Print("The request is served \n")
		c.Next()

	}

}
