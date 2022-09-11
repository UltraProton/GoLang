package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Reader! Your learning about 'goto' statement")
	// We create a for loop which runs until i is 10
	for i := 0; i < 10; i++ {
		fmt.Printf("Index: %d\n", i)

		switch i {
		case 5:
			goto exit
		case 6:
			goto second
		case 7: // I have to reference third label otherwise compiler wont run
			goto third
		}

	}

	fmt.Println("Skip this line here")
	// Create the exit label and insert code that should be executed when triggered
exit:
	fmt.Println("We are now exiting the program")

second:
	fmt.Println("This is a second label executed because it comes after exit")
	return

third:
	fmt.Println("this wont print")
}