package main


import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"net/url"
)

func main(){
	//PerformGetRequest("http://localhost:8000/get")
	//PerformPostJsonRequest("http://localhost:8000/post")
	PerformPostFormRequest("http://localhost:8000/postform")
}

func PerformGetRequest(myurl string) {

	response, err := http.Get(myurl)

	if err!= nil {
		panic(err)
	}

	defer response.Body.Close()

	var responseString strings.Builder //type of string which has lot of capabilities
	content, _ := ioutil.ReadAll(response.Body) //reading the response using ioutil its a raw response
	byteCount, _ := responseString.Write(content) //write the content in responseString and it returns bytecount
	fmt.Println(response.ContentLength, byteCount)
	fmt.Println(response)
	fmt.Println("Content: ",content)
	fmt.Println(string(content)) //simple way
	fmt.Println(responseString.String()) //using string builder
	
}

func PerformPostJsonRequest(myurl string){
	//create a request body using strings library
	requestBody := strings.NewReader(`
		{
			"coursename" : "aditya's course",
			"price": 100,
			"platform": "ultraproton.io"
		}
	`)

	response, err := http.Post(myurl,"application/json",requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}

func PerformPostFormRequest(myurl string) {
	//formdata
	data := url.Values{} //creating form
	data.Add("first", "aditya") //adding form values

	response, err := http.PostForm(myurl,data) //passing form data alongwith url

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body) //response body should always be read by ioutil and taken into content

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())

	fmt.Println(response)
} 