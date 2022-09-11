package main

import (
	"fmt"
	"net/url"
)

const urlA ="https://lco.dev:3000/learn?coursename=reactjs&payment=34lkajflkdf"

func main(){
	fmt.Println("Welcome to handling URLs in golang")
	 
	result, _ := url.Parse(urlA) //Parsing the url 

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery) 

	qparams := result.Query() //returns the query params as map

	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	//iterate the map

	for key, value := range qparams{
		fmt.Println(key,value)
	}

	//creating the url. Notice: The key values in this have to be exactly like written below
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}