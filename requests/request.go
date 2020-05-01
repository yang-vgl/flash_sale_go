package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"../actions"
	"../customers"
	"../products"
)

var Products struct {
	Number int
	Wg     *sync.WaitGroup
	List   chan int
}

func main() {

	ProductReady := products.GetProduct(500)

	var wg, productChan = products.PrepareProduct(ProductReady)

	ProductReady.Wg = wg
	ProductReady.List = productChan

	Products = ProductReady

	fmt.Println(Products)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      http.TimeoutHandler(http.HandlerFunc(purchaseRequest), 1*time.Second, "Sold Out!\n"),
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}

}

func purchaseRequest(w http.ResponseWriter, r *http.Request) {
	Products.Wg.Add(1)
	go actions.Buy(Products.List, Products.Wg, random(1, 20000))
	Products.Wg.Wait()
}

func slowHandler(w http.ResponseWriter, req *http.Request) {
	time.Sleep(2 * time.Second)
	println("slow")
	io.WriteString(w, "I am slow!\n")
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func getCustomers(list chan int) {
	time.Sleep(30 * time.Second)
	customers := customers.GetCustomers(list)
	for i, v := range customers {
		fmt.Println(i, v)
	}

}
