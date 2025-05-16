package service

import (
	"ServiceA/domain/entity"
	"ServiceA/infrastructure/dao"
	"github.com/gin-gonic/gin"

	"ServiceA/domain/repository"
	"ServiceA/infrastructure/model"
)

type AService struct {
	D repository.ItemRepository
}

func NewService(d *dao.Dao) *AService {
	return &AService{
		D: d,
	}
}
func (as *AService) Buy(c *gin.Context, id int, num int) (model.Response, error) {
	it, err := as.D.GetItem(id, c)
	if err != nil {
		return model.Response{}, err
	}

	// 调用 item 的 Buy 方法（领域逻辑）
	result := it.Buy(num)
	err = as.D.SetItem(c, it.Id, it)
	if err != nil {
		return model.Response{}, err
	}
	return result, nil

}
func (as *AService) Get(c *gin.Context, id int) (model.Response, error) {
	it, err := as.D.GetItem(id, c)
	if err != nil {
		return model.Response{}, err
	}
	return model.Response{
		Data: it,
	}, nil
}
func (as *AService) Set(c *gin.Context, it entity.Item) (model.Response, error) {
	err := as.D.SetItem(c, it.Id, it)
	if err != nil {
		return model.Response{}, err
	}
	return model.Response{
		Msg: "ok",
	}, nil
}
