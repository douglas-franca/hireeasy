package http

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, handler *Handler, auth gin.HandlerFunc) {
	routes := rg.Group("/companies", auth)
	routes.POST("", handler.Create)
	routes.GET("", handler.List)
	routes.GET("/:companyID", handler.Get)
}
