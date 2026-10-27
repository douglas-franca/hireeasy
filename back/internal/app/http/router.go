package http

import (
	"net/http"

	"github.com/douglas-franca/hireeasy/internal/platform/httperror"
	"github.com/gin-gonic/gin"
)

type healthResponse struct {
	Status string `json:"status"`
}

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger(), httperror.Recovery(), httperror.Middleware())
	r.NoRoute(func(c *gin.Context) {
		httperror.Abort(c, httperrorNotFound())
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, healthResponse{Status: "ok"})
	})

	registerSwagger(r)

	return r
}
