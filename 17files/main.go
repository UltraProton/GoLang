package main

import (
	"fmt"
	"io"
	"os"
	"io/ioutil"
)

func main() {
	fmt.Println("This is files demo")
	//x := 5 
	content := "This needs to go in a file"
	file, err := os.Create("./myfile.txt") //os might not create the file

	if err != nil {
		panic(err) //going to stop execution
	}
	
	length, err := io.WriteString(file , content) 

	//fmt.Println(x,content)
	fmt.Println("File length: ",length)
	
	defer file.Close() //recommended as you may be writing the code after this line also

	readFile("./myfile.txt")
}

func readFile(filename string){
	databyte, err := ioutil.ReadFile(filename)

	if err!= nil {
		panic(err)
	}

	fmt.Println("Text data inside the file is \n", string(databyte))
}