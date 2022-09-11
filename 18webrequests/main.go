package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const urlA ="https://lco.dev"
const urlB ="https://ultraproton.github.io"

func main() {
	response, err := http.Get(urlB)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type %T\n",response)
	fmt.Println(response)
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)


}