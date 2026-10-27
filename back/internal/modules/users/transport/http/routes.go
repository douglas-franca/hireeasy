package http

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, handler *Handler, auth gin.HandlerFunc) {
	userRoutes := rg.Group("/users")

	userRoutes.POST("", handler.Register)
	userRoutes.POST("/verification/confirm", handler.ConfirmVerification)

	me := userRoutes.Group("/me", auth)

	me.GET("", handler.GetCurrent)
	me.PATCH("", handler.UpdateCurrent)
	me.DELETE("", handler.DeleteCurrent)
	me.POST("/verification", handler.RequestVerification)
}
