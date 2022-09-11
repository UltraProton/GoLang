package main

import "fmt"

func main(){
	var fruit [4]string
	chocolates := [10]string{"cadbury", "nestle"} //slice
	fruit[0]="apple"
	fruit[1]="orange"
	fruit[3]="kiwi"
	fmt.Println("fruit list: ",fruit)

	fmt.Println("Length: ",len(fruit))
	fmt.Println("Chocolates: ", chocolates)

}