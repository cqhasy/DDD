package service

import (
	"ServiceB/domain/entity"
	"ServiceB/domain/repository"
	"ServiceB/infrastructure/dao"
	"context"
)

type BService struct {
	D repository.OrderRepository
}

func ProvideBRepository(d *dao.Dao) repository.OrderRepository {
	return d
}
func NewService(d repository.OrderRepository) *BService {
	return &BService{
		D: d,
	}
}
func (service *BService) Create(ctx context.Context, order entity.Order) error {
	err := service.D.MakeOrder(ctx, order)
	if err != nil {
		return err
	}
	return nil
}
