package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Orders struct which contains
// an array of orders
type Orders struct {
	Orders []Order `json:"orders"`
}

// Order struct which contains a name
// a type and a list of properties
type Order struct {
	Id				int `json:"id"`
	ProductId		int `json:"productId"`
	Properties		[]OrderProperties `json:"properties"`
	Quantity		int `json:"quantity"`
	OrderID			int	`json:"orderID"`
	Routes			Route `json:"route"`
}

type OrderProperties struct {
	Size	string  `json:"size"`
	Color	string  `json:"color"`
}

type Route struct {
	Driver	string `json:"driver"`
	Route	int `json:"route"`
}

func (m *Order) filterByRoute(r int) bool{
	var result bool
	if m.Routes.Route == r {
		result = true
	}
	return result
}

func (m *Order) filterByQuantity(q int) bool{
	var result bool
	if m.Quantity == q {
		result = true
	}
	return result
}

func (m *Order) filterByOrderID(o int) bool{
	var result bool
	if m.OrderID == o {
		result = true
	}
	return result
}

func GetOrders(w http.ResponseWriter, r *http.Request){
	var filterby FilterBy
	err := json.NewDecoder(r.Body).Decode(&filterby)
	if err != nil{
		log.Println("Error en consulta")
	}

	var order Orders
	var result Order
	order = ReadOrders()

	result = FilterOrders (order, filterby.Field, filterby.Value)
	w.Header().Set("Contend-Type", "application/json")
	j, err := json.Marshal(result)
	if err != nil{
		log.Println("Error en consulta")
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	OutputNote = nil
}

func ReadOrders() Orders{
	jsonFile, err := os.Open("orders.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened orders.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var orders Orders
	json.Unmarshal(byteValue, &orders)
	return orders
}


func FilterOrders (f Orders, filter string, value int) Order{
	var res bool
	var results Order
	for i := range f.Orders {
		switch filter{
		case "route":
			res = f.Orders[i].filterByRoute(value)
		case "orderID":
			res = f.Orders[i].filterByOrderID(value)
		case "quantity":
			res = f.Orders[i].filterByQuantity(value)
		}
		if res{
			fmt.Println(f.Orders[i])
			results.ProductId = f.Orders[i].ProductId
			results.Quantity = f.Orders[i].Quantity
			results.OrderID = f.Orders[i].OrderID
			results.Routes = f.Orders[i].Routes
		}
	}
	return results
}