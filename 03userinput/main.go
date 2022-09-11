package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to Mithai Shop"
	fmt.Println(welcome, "\n")
	fmt.Println("Enter the pizza rating")
	reader := bufio.NewReader(os.Stdin)

	/*comma ok
		IF the error comes it will be stored in "err" variable or if you don't care about that you can use 
		"_" for that.
	*/
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

}
