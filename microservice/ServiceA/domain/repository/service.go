package repository

import (
	"ServiceA/domain/entity"
	"context"
)

type ItemRepository interface {
	GetItem(id int, ctx context.Context) (entity.Item, error)
	SetItem(ctx context.Context, id int, item entity.Item) error
}
