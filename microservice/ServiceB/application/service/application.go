package service

import (
	"ServiceB/domain/entity"
	"ServiceB/domain/service"
	"context"
)

type Service interface {
	Create(ctx context.Context, order entity.Order) error
}
type Application struct {
	app Service
}

func ProvideService(se *service.BService) Service {
	return se
}
func NewApplication(app Service) *Application {
	return &Application{app: app}
}
func (a *Application) SetOrder(ctx context.Context, order entity.Order) error {
	err := a.app.Create(ctx, order)
	if err != nil {
		return err
	}
	return nil
}
