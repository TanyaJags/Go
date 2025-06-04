package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type cart []Item

// AddItem adds an item to the cart
func AddItem(Items []Item) cart {
	c := cart{}
	for _, item := range Items {
		c = append(c, item)
	}
	return c
}

func (c cart) RemoveItem(item Item) cart {
	var newCart cart
	for _, v := range c {
		if v.ID != item.ID {
			newCart = append(newCart, item)
		}
	}
	return newCart
}

func (c cart) ViewCart() {
	for _, item := range c {
		println("ID:", item.ID, "Name:", item.Name, "Price:", item.Price, "Qty:", item.Qty)
	}
}

func (c cart) saveToFileAsJSON(fileName string) error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0666)
}

func (c cart) searchItems(name string) []Item {
	var searchResults []Item
	for _, item := range c {
		if item.Name == name {
			searchResults = append(searchResults, item)
		}
	}
	return searchResults
}

func (c cart) getTotalPrice() float64 {
	var total float64
	for _, item := range c {
		total += item.Price * float64(item.Qty)
	}
	return total
}

func (c cart) emptyCart() cart {
	println("Cart has been emptied.")
	return cart{}
}

func (c cart) checkOut() cart {
	totalPrice := c.getTotalPrice()
	fmt.Printf("Total Price of items in the cart is: %.2f\n", totalPrice)
	c = c.emptyCart()
	fmt.Println("Cart has been checked out and is now empty.")
	return c
}
