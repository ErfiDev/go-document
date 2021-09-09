package main

import "fmt"

type human []string

func main() {
	humans := human{"erfan", "ali"}
	humans.update("john")

	fmt.Println(humans)
	// return ["erfan" , "ali"]
	// why????

	// Go is a pass by value language
	// what is the pass by value language?
	// pass by value language means that a copy
	// of the actual parameter's value made in memory

	// how a RAM works??
	// |address  |           value              |
	// | 0001    |            erfan             |
	// | 0002    |  []string{"erfan" , "ali"}   |

	// Each box in your computer's memory can store
	// some data and each one of these little boxes
	// there are these little value containers has
	// some discrate address

	// newHumans := human{"erfan"}
	// so when we create this newHumans variable of
	// type human go will create that slice of string
	// it will then go to the local memory on our
	// local machine
	// |   address   |   value   |
	// |   0001      |    ["erfan"]   |

	// what happen when we call this update("john") func?
	// whenever we pass some value into a function
	// will take that value its going to copy all of
	// that data inside that slice of string and then
	// place it inside of new some new container inside of
	// computer memory...
	// |   address   |        value          |
	// |   0001      |       ["erfan"]       |
	// |   0002      |  ["erfan" , "update"] |

	// we are actually not updating the original
	// slice of string.. we are updating the copy
	// how do we rectify this issue??
	// we should use Pointers to solve this problem

	// func (h *human) update(name string) {
	// h = append(h , name)
	// }

	// we solve the problem...

	// how we can find the pointer of variable?
	// x := "erfan"
	// pointerX := &x
	// pointerX variable contain the x variable
	// pointer address in memory

	// how we can receive pointer value in func?
	x := 1
	fmt.Println("initial: ", x)
	// return 1

	toZero(x)
	fmt.Println("to zero without using pointer: ", x)
	// return 1
	// becuase change the copy variable in memory

	toZeroWithPointer(&x)
	fmt.Println("to zero with pointer: ", x)
	// return 0
	// becuase change the original variable in memory

	// &Variable means give me the memory address
	// *Pointer means give me the value this memory
	// address is pointing at

	// Value types
	// Use pointers to change these things in a function
	// Int , String , Struct , Bool , Float

	// Refrence types
	// Don't worry about pointers with these
	// Slice , Maps , Channels , Pointers , Functions
}

func toZero(value int) {
	value = 0
}

func toZeroWithPointer(value *int) {
	*value = 0
}

func (h human) update(name string) {
	h = append(h, name)
}
