package dao

import (
	"ServiceB/domain/entity"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Dao struct {
	Redis *redis.Client
}

func NewDao(redis *redis.Client) *Dao {
	return &Dao{
		Redis: redis,
	}
}
func (dao *Dao) MakeOrder(ctx context.Context, order entity.Order) error {
	key := fmt.Sprintf("Order:%d", order.Id)
	data, err := json.Marshal(order)
	if err != nil {
		return err
	}
	err = dao.Redis.Set(ctx, key, data, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
