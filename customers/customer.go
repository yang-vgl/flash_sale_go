package customers

// GetCustomers : get all customers who successfully got the product
func GetCustomers(list chan string) []string {
	customers := make([]string, 0)
	for range list {
		customer := <-list
		customers = append(customers, customer)
	}
	return customers
}
