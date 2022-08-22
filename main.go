package main

import (
	"fmt"
	"log"

	"github.com/ykko21/go_package_1/helpers"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	randNum := helpers.RandomeNumber(numPool)
	intChan <- randNum
}

func main() {
	intChan := make(chan int)
	defer close(intChan)
	go CalculateValue(intChan)

	num := <-intChan //Listing for the channel
	log.Println(num)
}

func t2() { //channel test

}

func t1() {
	log.Println("Hello")
	var cust helpers.Customer
	cust.FirstName = "Peter"
	cust.CustomerId = "a3lsd0ee0dlcle9i"
	cust.CellPhoneNumber = "201-495-3399"
	fmt.Println(cust)
}
