package repository

import (
	"ServiceB/domain/entity"
	"context"
)

type OrderRepository interface {
	MakeOrder(ctx context.Context, order entity.Order) error
}
