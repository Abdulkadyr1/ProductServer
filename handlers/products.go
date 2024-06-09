package handlers

import (
	"fmt"
	"handlemod/data"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct{}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		product := data.GetProdList()
		//fmt.Println(product)
		product.ToJson(w)
	}
	if r.Method == http.MethodPost {
		product := data.ProductsType{}
		err := product.FromJson(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		data.AddToProdList(product)
		fmt.Printf("%s\n", product)
	}
	if r.Method == http.MethodPut {
		reg := regexp.MustCompile("/[0-9]+")
		g := reg.FindAllString(r.URL.Path, -1)

		if len(g) != 1 {
			http.Error(w, "Error handling url", http.StatusBadRequest)
		}
		if len(g[0]) != 2 {
			http.Error(w, "Error handling url", http.StatusBadRequest)
		}
		product := data.ProductsType{}
		err := product.FromJson(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		stringId := g[0][1]
		productList := data.GetProdList()
		id, err := strconv.Atoi(string(stringId))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		productList[id-1] = product[0]
		fmt.Println("ID: ", id)
		fmt.Printf("Products: %s", productList)
	}
	if r.Method == http.MethodDelete {
		fmt.Println("DELETE HANDLE")
		reg := regexp.MustCompile("/[0-9]+")
		g := reg.FindAllString(r.URL.Path, -1)
		if len(g) != 1 {
			http.Error(w, "Error handling url", http.StatusBadRequest)
		}
		stringid := g[0][1]
		intId, err := strconv.Atoi(string(stringid))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		data.DeleteFromProdList(intId - 1)
	}
}
