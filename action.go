package main

import (
	"./products"
	"fmt"
	"sync"
)

func Buy(list chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		val, ok := <-list
		if ok == false {
			fmt.Println("sold out !")
		}else{
			fmt.Println(val, ok)
		}
}

func main() {

	product := products.GetProduct(10)

	var wg, productChan = products.PrepareProduct(product)

	fmt.Println(len(productChan), cap(productChan))

	wg.Add(1)
	go Buy(productChan, wg)
	wg.Wait()
	close(productChan)

}
