package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("Welcome to mithai shop")
	fmt.Println("Please rate the pizza between 1 and 5")

	/*
	reader is initialized
	*/
	reader   := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') //take the input from the reader until delimiter value newline is processed

	fmt.Println("Thanks for rating, ",input)
	
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("We changed your rating: ", numRating+1)
		numRating+=1
		fmt.Println("numRating now: ", numRating) 
	}

}