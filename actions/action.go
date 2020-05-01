package actions

import (
	"fmt"
	"sync"

	"../products"
)

var Test string

func Buy(list chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	val, ok := <-list
	if ok == false {
		fmt.Println("sold out !")
	} else {
		fmt.Println(val, ok)
	}
}

func Test1() {
	fmt.Println("test")
}

func main() {

	Test = "test"

	ProductReady := products.GetProduct(10)

	var wg, productChan = products.PrepareProduct(ProductReady)

	ProductReady.Wg = wg
	ProductReady.List = productChan

	fmt.Println(ProductReady)
	/*
		wg.Add(1)
		go Buy(productChan, wg)
		wg.Wait()
		close(productChan)
	*/
}
