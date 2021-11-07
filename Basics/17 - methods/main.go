package main

import "fmt"

// Go does not have classes. However, we can define
// methods on types
type users []string

func main() {
	newUsers := users{"erfan", "john", "eminem", "mark"}

	// fmt.Println(newUsers)
	newUsers.logger()
	newUsers.custumLogger(2, 4)

	addedUsers := newUsers.addUser("erik")
	addedUsers.logger()
	addedUsers.length()
}

// A method is a function with a special receiver argument.
func (u users) logger() {
	fmt.Println(u)
}

func (u users) custumLogger(from, to int) {
	fmt.Println(u[from:to])
}

func (u users) addUser(name string) users {
	list := u
	list = append(list, name)

	return list
}

func (u users) length() {
	fmt.Println(len(u))
}
