package main

import (
	"fmt"
	"strings"
)

func main() {
	// string are basic data structure, and every language has a number
	// of predefined functions for manipulating strings
	// in GO these are gathered in package `strings`

	myName := "Erfan"

	// HasPrefix function tests whether the string s begins with prefix
	fmt.Println(strings.HasPrefix(myName , "Erf"))

	// Contains function testing whether a string contain a substring
	fmt.Println(strings.Contains(myName , "f"))

	// Index and LastIndex functions
	fmt.Println(strings.Index(myName , "n"))
	fmt.Println(strings.LastIndex(myName , "v"))

	// Replacing string
	fmt.Println(strings.Replace(myName , "fan" , "nan" , 2))

	// Count function:  return number of non-overlapping
	forTest := "Erfaaaaaaaaan"
	fmt.Println(strings.Count(forTest , "a"))

	// Repeat function
	fmt.Println(strings.Repeat("ba" , 2))

	// Changing the case of the string
	fmt.Println(strings.ToLower(myName) , strings.ToUpper(myName))

	// Trim functions
	withSpace := "Er fa n"
	trimLeftStr := "johnjohn";
	fmt.Println(strings.Trim(myName , "Er"))
	fmt.Println(strings.TrimSpace(withSpace))
	fmt.Println(strings.TrimLeft(trimLeftStr , "jo"))
	fmt.Println(strings.TrimRight(trimLeftStr , "hn"))
}