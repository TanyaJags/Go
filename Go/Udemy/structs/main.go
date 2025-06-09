package main

import "fmt"

type ContactInfo struct {
	email   string
	zipcode int
}

type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

func main() {
	// p := Person{
	// 	firstName: "John",
	// 	lastName:  "Doe",
	// }

	// println("First Name:", p.firstName)
	// println("Last Name:", p.lastName)

	// //another way to initialize a struct
	// p2 := Person{"Jane", "Smith"}
	// println("First Name:", p2.firstName)
	// println("Last Name:", p2.lastName) //this way is quite problematic because it is not clear which field is which, so it is better to use named fields as in the first example.
	// fmt.Println(p2)

	//Third way of doing things
	// var alex Person
	// alex.firstName = "Alex"
	// alex.lastName = "Johnson"
	// fmt.Println(alex)
	// fmt.Printf("First Name: %s, Last Name: %s\n", alex.firstName, alex.lastName)
	// fmt.Printf("%+v\n", alex) // %+v will print the struct with field names, which is useful for debugging

	//new person with contact infro
	jim := Person{
		firstName: "Jim",
		lastName:  "Halpert",
		contact: ContactInfo{
			email:   "JimHalpert@gamil.com",
			zipcode: 12345,
		},
	}
	fmt.Println(jim)
}
