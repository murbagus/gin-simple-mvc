package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthGlobalMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Melewati middleware auth global")
	}
}
