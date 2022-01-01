package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("start ozearth")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world ozearth !!")
}
