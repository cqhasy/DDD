package dao

import (
	"ServiceA/domain/entity"
	"context"
	"encoding/json"
	"errors"
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
func (dao *Dao) GetItem(id int, ctx context.Context) (entity.Item, error) {
	key := fmt.Sprintf("item:id:%d", id)
	data, err := dao.Redis.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return entity.Item{}, err
	} else if err != nil {
		return entity.Item{}, err
	}
	var item entity.Item
	err = json.Unmarshal([]byte(data), &item)
	if err != nil {
		return entity.Item{}, err
	}
	return item, nil
}
func (dao *Dao) SetItem(ctx context.Context, id int, item entity.Item) error {
	key := fmt.Sprintf("item:id:%d", id)
	data, err := json.Marshal(item)
	if err != nil {
		return err
	}
	err = dao.Redis.Set(ctx, key, data, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
