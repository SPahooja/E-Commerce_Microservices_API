package handlers

import (
	"GoMicroservice/data"
	"log"
	"net/http"
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
		p.addProduct(rw, r)
		return
	}
	if r.Method == http.MethodPut {
		path := r.URL.Path
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(path, -1)
		if len(g) != 1 {
			http.Error(rw, "invalid URL", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			http.Error(rw, "invalid URL", http.StatusBadRequest)
			return
		}
		ids := g[0][1]
		id, err := strconv.Atoi(ids)
		if err != nil {
			http.Error(rw, "invalid URL", http.StatusBadRequest)
			return
		}
		p.updateProduct(id, rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET")
	lp := data.GetProduct()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Failed to write from json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Failed to POST from request", http.StatusInternalServerError)
	}
	data.AddProduct(prod)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Failed to POST from request", http.StatusInternalServerError)
	}
	err = data.UpdateProduct(id, prod)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return

	}
}
