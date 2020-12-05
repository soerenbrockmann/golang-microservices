package handlers

import (
	"log"
	"net/http"
	"sbrockmann.com/product-api/data"
)

type Products struct{
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p* Products) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		p.getProducts(rw, r)
		return 
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p* Products) getProducts(rw http.ResponseWriter, h *http.Request){
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unabale to marshal json", http.StatusInternalServerError)
	}
}