package main

import "fmt"

const (
	erfan = iota
	john
	robert
	leo
)

func main() {
	// Iota is a useful concept for creating incrementing
	// constants in Go. However, there are several areas
	// where iota may not be appropriate to use.
	// This article will cover several different ways in
	// which you can use iota, and tips on where to be
	// cautious with it's use.
	fmt.Println(erfan, john, robert, leo)

	// Skipping values
	const (
		zero int = iota
		one
		_
		three
	)

	fmt.Println(zero, one, three)

	// Advanced Iota Usage
	// Because of the way iota automatically increments,
	// we can use it to calculate more advanced scenarios.
	// For instance, if you have worked with bitmasks
	// Iota can be used to quickly create the correct
	// values by using the bit shift operator.
	const (
		read   = 1 << iota // 00000001 = 1
		write              // 00000010 = 2
		remove             // 00000100 = 4

		// admin will have all of the permissions
		admin = read | write | remove
	)
	fmt.Println(read, write, remove, admin)

	// Memory size
	const (
		_  = 1 << (iota * 10) // ignore the first value
		KB                    // decimal:       1024 -> binary 00000000000000000000010000000000
		MB                    // decimal:    1048576 -> binary 00000000000100000000000000000000
		GB                    // decimal: 1073741824 -> binary 01000000000000000000000000000000
	)
	fmt.Println(KB, MB, GB)
}
