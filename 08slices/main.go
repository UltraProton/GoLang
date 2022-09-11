package main
import "fmt"


func main(){
	var fruitList = []string{"Apple","Mango"} //you are not specifying the size of the slice
	fmt.Printf("Type of fruitlist: %T\n", fruitList)
	
	fruitList = append(fruitList, "Guava", "Kiwi")
	fmt.Println("FruitList: ", fruitList)

	fruitList=  append(fruitList[1:2]) //append is used to slice the slice array

	fmt.Println("FruitList: ", fruitList)

	var courses = []string {"react", "java", "javascript", "C"}
	fmt.Println(courses)

	//remove a value from the slice using index

	courses=  append(courses[:2], courses[3:]...)

	fmt.Println(courses)

}