package main

import (
	"fmt"
	"time"
)

func cashflow(money chan int) {
	//amount := <-money
	//fmt.Println("Received", amount, "$!")

	//-- Select auf mehreren Channels
	select {
	case amount := <-money:
		fmt.Println("Received", amount, "$!")

	case <-time.After(1 * time.Second):
		fmt.Println("Got nothing!!")
	}
}

func main() {
	money := make(chan int)
	go cashflow(money)

	//money <- 1000

	time.Sleep(2 * time.Second)
}
