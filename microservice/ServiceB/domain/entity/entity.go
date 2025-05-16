package entity

import "time"

type Order struct {
	Num    int
	Prices int
	Time   time.Time
	Id     int64
}

func (o *Order) MakeOrder(n int, p int) {
	o.Num = n
	o.Prices = p
	o.Time = time.Now()
	o.Id = time.Now().Unix()
}
