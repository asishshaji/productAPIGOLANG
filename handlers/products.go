package handlers

import (
	"log"
	"net/http"
	"pApi/data"
	"regexp"
	"strconv"
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

	if r.Method == http.MethodPost {
		p.addProducts(rw, r)
		return
	}

	if r.Method == http.MethodPut {

		reg := regexp.MustCompile(`/{[0-9]+}`)
		grp := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(grp) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(grp[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString := grp[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.updateProducts(rw, r, id)
		return
	}

}

func (p *Products) updateProducts(rw http.ResponseWriter, r *http.Request, id int) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err = data.UpdateProducts(id, prod)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (p *Products) addProducts(rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}

	data.AddProduct(prod)

}

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
