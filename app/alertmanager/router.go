package alertmanager

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.POST("/alertmanager", alertManagerHandler)
}
