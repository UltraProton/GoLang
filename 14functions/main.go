 package main

 import (
	"fmt"
 )

 /* Variable argument list is also a slice
 */
 func mulA(values ...int) (int, string){
	ans := 1
	for _, value := range values{
		ans*=value
	}

	return ans,"Hi there"
 }

 /* Passing a slice as arg
 */
 func mulB(values []int)(int, string){
	ans := 1
	for _, value := range values{
		ans*=value
	}

	return ans,"Hi there"
 }

 func main(){
	ansB, _ := mulB([]int{2,3,4,5})
	ansA,_ := mulA(2,3,4,5)
	fmt.Println(ansA,ansB)
 }