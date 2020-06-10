package middleware

import "github.com/gin-gonic/gin"

func CorsMiddle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		ctx.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
		ctx.Next()
	}
}
