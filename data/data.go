package data

import (
	"encoding/json"
	"io"
)

type Product struct {
	Id   int
	Name string
}

var prodList = ProductsType{
	&Product{
		Id:   1,
		Name: "Tea",
	},
	&Product{
		Id:   2,
		Name: "Coffe",
	},
}

func AddToProdList(a ProductsType) ProductsType {
	prodList = append(prodList, a...)
	return prodList
}

func (a *ProductsType) ToJson(f io.Writer) {
	ec := json.NewEncoder(f)
	ec.Encode(a)
}

func (a *ProductsType) FromJson(f io.Reader) error {
	ec := json.NewDecoder(f)
	return ec.Decode(a)
}

func GetProdList() ProductsType {
	return prodList
}

type ProductsType []*Product
