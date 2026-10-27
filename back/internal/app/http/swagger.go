package http

import (
	_ "embed"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

//go:embed openapi.json
var openAPI []byte

func registerSwagger(router *gin.Engine) {
	router.GET("/swagger/openapi.json", func(c *gin.Context) { c.Data(nethttp.StatusOK, "application/json", openAPI) })
	router.GET("/swagger/index.html", func(c *gin.Context) {
		c.Data(nethttp.StatusOK, "text/html; charset=utf-8", []byte(`<!doctype html><html><head><title>HireEasy API</title><link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist/swagger-ui.css"></head><body><div id="swagger-ui"></div><script src="https://unpkg.com/swagger-ui-dist/swagger-ui-bundle.js"></script><script>SwaggerUIBundle({url:'/swagger/openapi.json',dom_id:'#swagger-ui'})</script></body></html>`))
	})
}
