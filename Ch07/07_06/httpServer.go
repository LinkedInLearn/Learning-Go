package main

import (
	"fmt"
	"net/http"
)

type server struct{}

func (h server) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1> Hello from Go!</h1>")
}

func main() {
	var webServer server
	err := http.ListenAndServe("localhost:1337", webServer)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}