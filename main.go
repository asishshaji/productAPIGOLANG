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
		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(rw, "Errror reading data", http.StatusBadRequest)
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("Error reading data."))
			return
		}

		fmt.Fprintf(rw, "Hello %s", data)

	})

	http.ListenAndServe(":9092", nil)
}
