package main

import (
	"fmt"
	"goge/extras"
)

func main() {
	fmt.Println("Press any key: ")

	// This should work in any platform
	keyPressed := extras.GetKeyPress()

	switch keyPressed {
	case 10:
		fmt.Println("You pressed intro")
		break
	case 32:
		fmt.Println("You pressed space bar")
	default:
		fmt.Println("You pressed a key with ASCII code ", keyPressed)
	}
}
