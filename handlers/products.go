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
			fmt.Println(err)
		}
		data.AddToProdList(product)
		fmt.Printf("%s\n", product)
	}
	if r.Method == http.MethodPut {
		reg := regexp.MustCompile("/[0-9]+")
		g := reg.FindAllString(r.URL.Path, -1)

		if len(g) != 1 {
			fmt.Println("Error")
			return
		}
		if len(g[0]) != 2 {
			fmt.Println("Error")
			return
		}
		product := data.ProductsType{}
		err := product.FromJson(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		stringId := g[0][1]
		productList := data.GetProdList()
		id, err := strconv.Atoi(string(stringId))
		if err != nil {
			fmt.Println("Error")
		}
		productList[id-1] = product[0]
		fmt.Println("ID: ", id)
		fmt.Printf("Products: %s", productList)
	}
}
