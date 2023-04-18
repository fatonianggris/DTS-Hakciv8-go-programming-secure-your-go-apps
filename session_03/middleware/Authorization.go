package middleware

import (
	"go-programming-secure-your-go-apps/session_03/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizeRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, err := helper.VerifyToken(ctx)
		if err != nil {
			status := http.StatusUnauthorized
			ctx.AbortWithStatusJSON(status, helper.ResponseError(status))
			return
		}

		ctx.Set("userData", userData)
		ctx.Next()
	}
}
