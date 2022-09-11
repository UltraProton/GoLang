package main

import (
	"fmt"
)

func main(){
	x :=10
	y :=3
	var ptr1 *int = &x  //one way of initializing pointer
	ptr2 := &y //straight forward way of initializing pointer

	fmt.Println("Value of ptr1: ",*ptr1)
	fmt.Println("Value of ptr2: ",*ptr2)

}