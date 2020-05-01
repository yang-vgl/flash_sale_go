package orders

import "fmt"

var number = 1

// ProcessOrder : process order after purchase
func ProcessOrder(user int) {
	fmt.Printf("User ID %d has successfully purchased one product, now process order %d \n", user, number)
	number++
}
