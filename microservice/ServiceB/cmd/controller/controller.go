package controller

import (
	"ServiceB/application/service"
	"ServiceB/cmd/remote"
	"ServiceB/domain/entity"
	"ServiceB/infrastructure/model"
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Reply struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}
type Controller struct {
	application Application
}
type Application interface {
	SetOrder(ctx context.Context, order entity.Order) error
}

func ProvideApplication(app *service.Application) Application { return app }
func NewController(application Application) *Controller {
	return &Controller{
		application: application,
	}
}

// SetOrder @Summary 提交订单
// @Description 通过此接口提交一个订单，包含商品数量和价格
// @Tags Order
// @Accept json
// @Produce json
// @Param req body model.SetOrderReq true "订单请求参数"  // 请求体类型是 SetOrderReq
// @Success 200 {object} dto.Response "成功"
// @Failure 400 {object} dto.Response "请求参数错误"
// @Failure 500 {object} dto.Response "服务器错误"
// @Router /order [post]
func (c *Controller) SetOrder(ctx *gin.Context, req model.SetOrderReq) (Reply, error) {
	var order entity.Order
	order.MakeOrder(req.Num, req.Price)
	err := c.application.SetOrder(ctx, order)
	if err != nil {
		return Reply{
			Data: nil,
			Msg:  err.Error(),
		}, err
	}
	path := remote.BuyRemote()
	id := strconv.Itoa(int(req.Id))
	num := strconv.Itoa(int(req.Num))
	data, err := remote.GetOtherServiceWithParams(path, id, num)
	if err != nil {
		return Reply{
			Data: nil,
			Msg:  err.Error(),
		}, err
	}
	return Reply{
		Data: string(data),
		Msg:  "下单成功",
	}, nil
}
