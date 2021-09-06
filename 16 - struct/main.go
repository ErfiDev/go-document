package main

import "fmt"

// Data structure. collection of properties that are
// related together
// This car struct type has name, model,
// engine and price fields
type car struct {
	name   string
	model  string
	engine string
	price  string
}

func main() {
	// New car
	var myCar car

	// Set values
	myCar.engine = "hemi 6.2L supercharged"
	myCar.model = "dodge"
	myCar.name = "hellcat redeye 2021"
	myCar.price = "84,000$"

	fmt.Println(myCar.name)

	// or we can declare and set values with different
	// way...
	yourCar := car{
		"M7",
		"BMW",
		"5.4L twin turbo 550HP 682NM",
		"120,000$",
	}

	fmt.Println(yourCar)
}
