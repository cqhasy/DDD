package application

import (
	"ServiceA/domain/entity"
	"ServiceA/domain/service"
	"ServiceA/infrastructure/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Application struct {
	S1 *service.AService
}

func NewApplication(aService *service.AService) *Application {
	return &Application{
		S1: aService,
	}
}
func (app *Application) BuyHandler(c *gin.Context) {

	idStr := c.Param("id")

	id, _ := strconv.Atoi(idStr)
	numStr := c.Param("num")
	num, err := strconv.Atoi(numStr)

	if err != nil || num <= 0 {
		c.JSON(http.StatusBadRequest, model.Response{
			Msg: "无效的购买数量",
		})
		return
	}
	re, err := app.S1.Buy(c, id, num)
	if err != nil {
		c.JSON(400, model.Response{
			Msg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, re)

}
func (app *Application) GetInfoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	re, err := app.S1.Get(c, id)
	if err != nil {
		c.JSON(200, model.Response{
			Msg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, re)
}

// AddHandler handles the adding of an item to the system.
// @Summary Add a new item
// @Description This endpoint allows you to add a new item to the inventory.
// @Tags items
// @Accept  json
// @Produce  json
// @Param item body entity.Item true "Item details"
// @Success 200 {object} model.Response "Successfully added the item"
// @Failure 400 {object} model.Response "Invalid input data"
// @Router /items [post]
func (app *Application) AddHandler(c *gin.Context) {
	var it entity.Item
	err := c.ShouldBind(&it)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Msg: err.Error(),
		})
		return
	}
	re, err := app.S1.Set(c, it)
	if err != nil {
		c.JSON(400, model.Response{
			Msg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, re)
}
