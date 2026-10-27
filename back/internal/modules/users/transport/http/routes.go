package http

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, handler *Handler) {
	userRoutes := rg.Group("/users")

	userRoutes.POST("", handler.CreateUser)
}
