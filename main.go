package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hi")
		data, _ := ioutil.ReadAll(r.Body)

		fmt.Fprintf(rw, "Hello %s", data)

	})

	http.ListenAndServe(":9092", nil)
}
