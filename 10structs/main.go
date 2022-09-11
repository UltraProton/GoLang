package main

import "fmt"

func main(){
	aditya := User{"Aditya","aditya@google.com","100"}
	fmt.Println(aditya)
	
	//+v gives key value of the struct also
	fmt.Printf("User: %+v\n",aditya)
}

type User struct {
	Name string 
	Email string
	ID string
}