package main

import "fmt"

func LinearSearch() {
	numberList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(numberList); i++ {
		if numberList[i] == 5 {
			fmt.Printf("Found 5 at %d\n", i+1)
			break
		}
	}
	fmt.Println("Linear Search Over")
}
