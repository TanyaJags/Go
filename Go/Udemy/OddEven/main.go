package main

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, number := range numbers {
		printoddoreven(number)
	}
	println("=End of Program=")
}

func printoddoreven(number int) {
	if number%2 == 0 {
		println(number, "is even")
	} else {
		println(number, "is odd")
	}
}
