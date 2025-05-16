package ginx

import (
	"ServiceB/application/dto"
	"ServiceB/cmd/controller"
	"github.com/gin-gonic/gin"
)

func TurnHandlerWithReq[Req any](fn func(*gin.Context, Req) (controller.Reply, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Req
		err := dto.Bind(c, &req)
		if err != nil {
			dto.Respond(c, nil, err, "")
			return
		}
		resp, err := fn(c, req)
		if err != nil {
			dto.Respond(c, c, err, "")
			return
		}
		dto.Respond(c, resp.Data, err, resp.Msg)
	}
}
