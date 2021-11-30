package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// The Go programming language supports several
	// bitwise operators including the followings:

	// &   bitwise AND
	// |   bitwise OR
	// ^   bitwise XOR
	// &^   AND NOT
	// <<   left shift
	// >>   right shift

	// &
	// In Go, the & operator performs the bitwise AND
	// operation between two integer operands
	// The AND operator has the nice side effect of
	// selectively clearing bits of an integer value to
	// zero. For instance, we can use the & operator to
	// clear (set to zero) the last 4 least significant
	// bits (LSB) to all zeros.

	var x uint8 = 0xAC // x = 10101100
	x = x & 0xF0       // x = 10100000
	// or x &= 0xF0

	// Another neat trick you can do with & operator is
	// to test whether a number is odd or even. This works
	// because a number is odd when its least significant
	// bit is set (equal 1). We can use the & operator
	// apply a bitwise AND operation to an integer the
	// value 1. If the result is 1, then the original
	// number is odd.

	for x := 0; x < 100; x++ {
		num := rand.Int()
		if num&1 == 1 {
			fmt.Printf("%d is odd\n", num)
		} else {
			fmt.Printf("%d is even\n", num)
		}
	}

	var odd int = 7
	var even int = 8
	fmt.Println("example: how it work:", odd&1)
	// if return 1 value is odd
	// and if return 0 value is even
	fmt.Println("example: how it work:", even&1)

	// |
	var value1, value2 int = 5, 3
	result := value1 | value2
	fmt.Println("5 | 3 = ", result)
	// Takes two numbers as operands and does OR on every
	// bit of two numbers. The result of OR is 1 any of
	// the two bits is 1.

	// ^
	// Takes two numbers as operands and does XOR on every
	// bit of two numbers. The result of XOR is 1 if the
	// two bits are different.
	result = value1 ^ value2
	fmt.Println("5 ^ 3 = ", result)

	// <<
	// Takes two numbers, left shifts the bits of the
	// first operand, the second operand decides the
	// number of places to shift.
	result = 5 << 3
	// 5 << 1  ===  10
	// 5 << 2  ===  20
	// 5 << 3  ===  40
	fmt.Println("5 << 3 = ", result)

	// >>
	// Takes two numbers, right shifts the bits of the
	// first operand, the second operand decides the number
	// of places to shift.
	result = 5 >> 1
	fmt.Println("5 >> 1 = ", result)

	// &^
	// This is a bit clear operator.
	result = 5 &^ 1
	fmt.Println("5 &^ 1 = ", result)
}
