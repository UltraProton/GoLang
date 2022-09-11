package main

import "fmt"

const global =" kumar"
const myVar = "aljdfkas"

func main() {
	var username string = "Aditya"
	fmt.Println(username)
	fmt.Printf("Variable type: %T \n",username)
	var id int =2
	fmt.Printf("Variable type is : %T",id)

	//implicit type
	var website = "ultraproton.github.io"
	fmt.Println("website: ",website)
	
	//walrus operator
	pass := "lakjfelsakdf"
	fmt.Println("pass: ",pass)
	fmt.Printf("%T",pass)
	fmt.Println(global)
}