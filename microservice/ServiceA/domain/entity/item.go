package entity

import "ServiceA/infrastructure/model"

type Item struct {
	Id    int
	Name  string
	Price int
	Num   int
}
type Bill struct {
	Num    int
	Prices int
}

func (item *Item) Buy(num int) model.Response {
	if num > item.Num {
		return model.Response{
			Msg:  "库存不足",
			Data: nil,
		}
	}
	item.Num = item.Num - num
	var b Bill
	b.Num = num
	b.Prices = item.Price * num

	return model.Response{
		Msg:  "success",
		Data: b,
	}
}
func (item *Item) GetInfo() model.Response {
	return model.Response{
		Msg:  "success",
		Data: item,
	}
}
