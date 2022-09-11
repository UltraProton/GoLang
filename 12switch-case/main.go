package main

import "fmt"

func main(){
	fmt.Println("Switch in go lang")
	//rand.Seed(time.Now().UnixNano())

	//dice := rand.Intn(6)+1

	dice := 4
	switch dice{
	case 1: fmt.Println("Value is 1")
	case 2: fmt.Println("Value is 2")
	default: fmt.Println("Don't know this value")
	}
}