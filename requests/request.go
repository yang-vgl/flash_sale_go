package main

import (
	"fmt"
	"net/http"
	"../products"
	"../actions"
	"sync"
)

var Products struct {
	Number int
	Wg *sync.WaitGroup
	List chan int
}

func main() {

	ProductReady := products.GetProduct(10)

	var wg, productChan = products.PrepareProduct(ProductReady)

	ProductReady.Wg = wg
	ProductReady.List = productChan

	Products = ProductReady

	fmt.Println(Products)

	http.HandleFunc("/purchase", purchaseRequest)
	http.ListenAndServe(":8080", nil)

}


func purchaseRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println(Products)
	Products.Wg.Add(1)
	go actions.Buy(Products.List, Products.Wg)
	Products.Wg.Wait()
}
