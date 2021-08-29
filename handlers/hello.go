package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hi")
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Errror reading data", http.StatusBadRequest)
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("Error reading data."))
		return
	}

	fmt.Fprintf(rw, "Hello %s", data)
}
