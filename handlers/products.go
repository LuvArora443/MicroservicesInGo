package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/LuvArora443/MicroservicesInGo/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products{
	return &Products{l}
}


func(p *Products) GetProducts(w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle GET method")
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err!=nil{
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST method")
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(prod)
}

func (p *Products) UpdateProducts(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err!=nil{
		http.Error(w, "Unable to convert id to string", http.StatusBadRequest)
		return
	}
	p.l.Println("Handle PUT method", id)
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	err = data.UpdateProduct(id, prod)
	if err==data.ErrProdNotFound{
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err!=nil{
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return 
	}

}

type KeyProduct struct{}
func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		prod := &data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}
	ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
	req := r.WithContext(ctx)
	next.ServeHTTP(w, req)
})
}