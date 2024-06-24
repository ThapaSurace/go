package main

import (
	"fmt"
)

func main() {
	myCar := struct {
		Make  string
		Model string
	}{
		Make:  "tesla",
		Model: "model 3",
	}

	fmt.Println(myCar)

	//You can even nest anonymous structs as fields within other structs
	type car struct {
		Make   string
		Model  string
		Height int
		Width  int
		// Wheel is a field containing an anonymous struct
		Wheel struct {
			Radius   int
			Material string
		}
	}

}
