package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"io/ioutil"
	"strings"
)


type Tour struct {
	Name string
	Price string
}

func main() {

	//Set the url link
	url := "http://services.explorecalifornia.org/json/tours.php"

	//get the page data returns as string
	content := contentFromServer(url)

	//fmt.Println(content)
	println("**********************************************")
	tours := toursFromJSON(content)
	//fmt.Println(tours)

	for _, tour := range tours {
		price, _, _ := big.ParseFloat(tour.Price, 10, 2, big.ToZero)
		//tour.Price = price
		fmt.Printf("%v (%.2f)\n", tour.Name, price)
	}
}

/**
Check if there's an error
If error is not null
then there is an error
 */
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func contentFromServer(url string) string {

	//get data from page/url and returns data | can return error
	resp, err := http.Get(url)
	//check if we got error from api call if so stop
	checkError(err)
	//close the api call after everything else done running
	defer resp.Body.Close()
	//read the parse the resposne type into bytes
	bytes, err := ioutil.ReadAll(resp.Body)
	//see if it got an error if so stop
	checkError(err)
	//return the bytes turned into string
	return string(bytes)
}

func toursFromJSON(content string) []Tour {
	//Makes a slice of the Tour type with starting size 0 but max size 20
	tours := make([]Tour, 0, 20)
	//stores JSON info ?
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)

	var tour Tour

	//loop thru json
	for decoder.More() {
		//get object, put in tour
		err := decoder.Decode(&tour)
		checkError(err)
		//assign tours = tours + tour object
		tours = append(tours, tour)
	}

	return tours
}