package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {
	fmt.Println("Start the service")
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
		return
	}

}
