package products

import "sync"

type Products struct {
	Number int
	Wg     *sync.WaitGroup
	List   chan int
}

type Test struct {
	Number int
	Wg     *sync.WaitGroup
	List   chan int
}

func GetProduct(number int) Products {
	products := Products{
		Number: number,
	}
	return products
}

func PrepareProduct(product Products) (*sync.WaitGroup, chan int) {
	wg := new(sync.WaitGroup)

	list := make(chan int, product.Number)

	for i := 0; i < product.Number; i++ {
		list <- i
	}
	return wg, list
}
