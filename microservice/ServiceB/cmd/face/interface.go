package face

import (
	"ServiceB/cmd/controller"
	"ServiceB/cmd/ginx"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(c *controller.Controller) *gin.Engine {
	r := gin.New()
	r.POST("/Order", ginx.TurnHandlerWithReq(c.SetOrder))
	return r
}
