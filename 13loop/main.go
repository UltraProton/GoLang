package main

import (
	"fmt"
)


func main(){
	cities := []string {"mzn","meerut"}

	fmt.Printf("%v\n",cities)

	for i:=0; i< len(cities); i++ {
		fmt.Println(cities[i])
	}

	//here i gives the index and not the actual value
	for i := range cities {
		fmt.Println(cities[i])
	}

	//kind of for each loop
	for index, value :=range cities {
		fmt.Printf("index is %v and value is %v\n",index,value)
	}

	Value :=1

	for Value < 10{
		fmt.Println("Value is ",Value)
		if Value == 5{
			break
		}
		if Value ==3{
			Value++
			continue
		}

		if Value ==11 {
			fmt.Println("We saw a value of",Value )
			goto come
		}
		Value++
		/*Imp if you write the label in the main then it will be executed like a normal statement
		in the regular flow */ 
		come:
			fmt.Println("From goto Value is :",Value)
		
	}

	fmt.Println(Value)
	
	/*
	come:
		fmt.Println("From goto Value is :",Value)
	*/
	
}