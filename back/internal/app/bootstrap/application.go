package bootstrap

import "github.com/gin-gonic/gin"

type App struct {
	Router *gin.Engine
	close  func()
}

func (a *App) Run(address string) error { return a.Router.Run(address) }
func (a *App) Close() {
	if a.close != nil {
		a.close()
	}
}
