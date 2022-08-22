package main

import (
	"fmt"
	"log"

	"github.com/ykko21/go_package_1/helpers"
)

func main() {
	log.Println("Hello")
	var cust helpers.Customer
	cust.FirstName = "Peter"
	cust.CustomerId = "a3lsd0ee0dlcle9i"
	cust.CellPhoneNumber = "201-495-3399"
	fmt.Println(cust)
}
