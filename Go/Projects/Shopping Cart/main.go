package main

import "fmt"

func main() {
	fmt.Println("This is my FIRST independent Go Project - SHOPPING CART")

	//TEST STRUCT ITEM - PASSED
	// item := Item{
	// 	ID:    1,
	// 	Name:  "Full Cream Milk",
	// 	Price: 34.0,
	// 	Qty:   1,
	// }

	// fmt.Printf("ID: %d  Name: %s  Price: %.2f  Quantity: %d\n",
	// 	item.ID, item.Name, item.Price, item.Qty)

	//Adding New Items in inventory may be
	items := []Item{
		NewItem(1, "Full Cream Milk", 34.0, 1),
		NewItem(2, "Bread", 20.0, 2),
		NewItem(3, "Eggs", 50.0, 1),
		NewItem(4, "Butter", 25.0, 1),
		NewItem(5, "Cheese", 40.0, 1),
		NewItem(6, "Yogurt", 30.0, 1),
		NewItem(7, "Chicken", 150.0, 1),
		NewItem(8, "Fish", 200.0, 1),
		NewItem(9, "Rice", 60.0, 1),
	}

	// fmt.Println("Items in Inventory:")
	// for _, item := range items {
	// 	fmt.Printf("ID: %d  Name: %s  Price: %.2f  Quantity: %d\n",
	// 		item.ID, item.Name, item.Price, item.Qty)
	// }

	// Adding items to the cart
	YourCart := cart{}
	YourCart = AddItem(items)

	// Viewing the cart
	fmt.Println("\nItems in Your Cart:")
	//YourCart.ViewCart()

	//Save the cart to a file
	err := YourCart.saveToFileAsJSON("cart.json")
	if err != nil {
		fmt.Println("Error saving cart to file:", err)
	} else {
		fmt.Println("Cart saved successfully to cart.json")
	}

	fmt.Print("\nSearching for 'Bread' in your cart...\n")
	searchResults := YourCart.searchItems("Bread")
	fmt.Println("Search Results:", searchResults)

	fmt.Println("Checking out the cart...")
	YourCart = YourCart.checkOut()

	//viewing the cart
	YourCart.ViewCart()
	fmt.Println("End of the program")
}
