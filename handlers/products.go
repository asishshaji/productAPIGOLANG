package handlers

import (
	"log"
	"net/http"
	"pApi/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		p.updateProducts(rw, r)
		return
	}

}

func (p *Products) updateProducts(rw http.ResponseWriter, r *http.Request) {}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	// s, err := json.Marshal(lp)
	// if err != nil {
	// 	http.Error(rw, "Error getting data", http.StatusInternalServerError)
	// }
	// rw.Write(s)

	// below method dont need any variables to store the data
	// if the data to be loaded is very large, then we can use
	// this to increase performance
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}
