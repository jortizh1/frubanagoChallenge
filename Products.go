package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
)

// Products struct which contains
// an array of products
type Products struct {
	Products []Product `json:"products"`
}

// Product struct which contains a name
// a type and a list of properties
type Product struct {
	Id				int `json:"id"`
	Category		string `json:"category"`
	Owner_name		string `json:"owner_name"`
	Product_name	string `json:"product_name"`
	Name string		`json:"name"`
	Properties		[]Properties `json:"properties"`
	Weight int		`json:"weight"`
	Price int		`json:"price"`
}

type Properties struct {
	Name string `json:"name"`
}

type ByName []Product
func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

type ByWeight []Product
func (a ByWeight) Len() int           { return len(a) }
func (a ByWeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByWeight) Less(i, j int) bool { return a[i].Weight < a[j].Weight }

func GetProducts(w http.ResponseWriter, r *http.Request){
	var orderby OrderBy
	err := json.NewDecoder(r.Body).Decode(&orderby)
	if err != nil{
		log.Println("Error en consulta")
	}

	var product Products
	product = ReadProducts()

	switch orderby.Field {
	case "Name":
		sort.Sort(ByName(product.Products))
	case "Weight":
		sort.Sort(ByWeight(product.Products))
	default:
		sort.Sort(ByName(product.Products))
	}
	w.Header().Set("Contend-Type", "application/json")
	j, err := json.Marshal(product)
	if err != nil{
		log.Println("Error en consulta")
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	OutputNote = nil
	OutputNote = make(map[string]Products)
}

func ReadProducts() Products{
	jsonFile, err := os.Open("products.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened products.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var product Products
	json.Unmarshal(byteValue, &product)
	return product
}