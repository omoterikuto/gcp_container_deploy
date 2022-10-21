package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", test)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "container deploy successed!")
}
