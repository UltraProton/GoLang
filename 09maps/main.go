package main

import (
	"fmt"
)

func main(){

	languages := make(map[string] string)

	languages["JS"]= "JAVASCRIPT"
	languages["PY"]= "PYTHON"

	fmt.Println(languages["JS"])

	delete(languages,"JS")

	fmt.Println(languages["PY"])

	//loops in go lang
	for _ , value := range languages{
		fmt.Println("value is ",value)
	}
}