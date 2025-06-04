package main

type Item struct {
	ID    int
	Name  string
	Price float64
	Qty   int
}

func NewItem(id int, name string, price float64, qty int) Item {
	return Item{
		ID:    id,
		Name:  name,
		Price: price,
		Qty:   qty,
	}
}
