package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"github.com/LuvArora443/MicroservicesInGo/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products{
	return &Products{l}
}

func(p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if r.Method==http.MethodGet{
		p.getProducts(w, r)
		return 
	}
	if r.Method==http.MethodPost{
		p.addProduct(w,r)
		return
	}
	
	if r.Method==http.MethodPut{
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g)!=1{
			http.Error(w, "URI not valid", http.StatusBadRequest)
			return
		}
		if len(g[0])!=2{
			http.Error(w, "URI not valid", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err!=nil{
			http.Error(w, "URI not valid", http.StatusBadRequest)
		}
		p.updateProducts(id, w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func(p *Products) getProducts(w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle GET method")
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err!=nil{
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST method")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}
	data.AddProduct(prod)
}

func (p *Products) updateProducts(id int, w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle PUT method")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}
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
