package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "http://services.explorecalifornia.org/json/tours.php"

	response, err := http.Get(url)

	checkError(err)
	fmt.Printf("Type %T", response)

	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	checkError(err)

	content := string(bytes)
	println(content)
}

func checkError(err error){
	if err != nil {
		panic(err)
	}
}