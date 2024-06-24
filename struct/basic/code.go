package main

import (
	"fmt"
)

//struct - collection of key, value  pairs

type user struct {
	fname    string
	lname    string
	age      int
	username string
	password string
}

func userDetails(u user) {
	fmt.Printf("First Name: %v\n", u.fname)

	fmt.Printf("Last Name: %s\n", u.lname)
}

func main() {
	userInfo := user{
		"surace", "thapa", 24, "surace", "surace",
	}

	fmt.Printf("First Name: %v\n", userInfo.fname)
	fmt.Printf("Last Name: %s\n", userInfo.lname)

	//passing value to function
	userDetails(userInfo)
}
