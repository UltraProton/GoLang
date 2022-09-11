package main

import "fmt"

func main(){
	aditya := User{"Aditya","aditya@google.com","100","25",true}
	fmt.Println(aditya)
	
	//+v gives key value of the struct also
	fmt.Printf("User: %+v\n",aditya)

	msg := aditya.NewMail("AditVP@google.com")

	fmt.Println(msg)
	fmt.Println(aditya)
}

type User struct {
	Name string 
	Email string
	ID string
	oneAge string
	Status bool
}

/* This method is associated with the struct User
*/
func (u *User) NewMail(email string) string {
	u.Email =email
	fmt.Println("Email is :", u.Email)

	return "Hi"
}