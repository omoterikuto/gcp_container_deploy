package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", test)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("start server")
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "container deploy is successed!! re")
}
