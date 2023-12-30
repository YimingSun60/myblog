package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	err := HomeHandler(w)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
		return
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the About page and 2 + 2 is %d", sum))
}

// 第一个字母小写的函数是private
func addValues(x, y int) int {
	sum := x + y
	return sum
}

func HomeHandler(w http.ResponseWriter) error {
	_, err := fmt.Fprintf(w, "This is the Home page")
	return err
}

func main() {
	fmt.Println("Start the service")
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
		return
	}

}
