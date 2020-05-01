package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"../actions"
	"../products"
)

var Products struct {
	Number int
	Wg     *sync.WaitGroup
	List   chan int
}

func main() {

	ProductReady := products.GetProduct(5)

	var wg, productChan = products.PrepareProduct(ProductReady)

	ProductReady.Wg = wg
	ProductReady.List = productChan

	Products = ProductReady

	fmt.Println(Products)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      http.TimeoutHandler(http.HandlerFunc(purchaseRequest), 1*time.Second, "Timeout!\n"),
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}

func purchaseRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println(Products)
	Products.Wg.Add(1)
	go actions.Buy(Products.List, Products.Wg)
	Products.Wg.Wait()
}

func slowHandler(w http.ResponseWriter, req *http.Request) {
	time.Sleep(2 * time.Second)
	println("slow")
	io.WriteString(w, "I am slow!\n")
}
