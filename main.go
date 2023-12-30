package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the About page and 2 + 2 is %d", sum))
}

//第一个字母小写的函数是private
func addValues(x, y int) int {
	sum := x + y
	return sum
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
		return
	}

}
