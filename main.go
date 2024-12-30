package main

import (
	"fmt"
	"net/http"
)

var check string

func init() {

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "website test 123")
}

func main() {
	fmt.Println("test")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
