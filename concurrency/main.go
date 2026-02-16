package main

import "fmt"

// Concurrency
// Goroutines --> leichtgewichtiger Thread der Go-Runtime
// Channels --> Kanal für Nachrichten mit den Aufgaben: 1 Kommunikation und 2 Synchronisation von Goroutinen

func main() {
	//-- Gorountinen und Channels

	// 1. Channel erzeugen
	money := make(chan int)

	go func() {
		// 3. Receive auf Channel money
		amount := <-money // Wartet solange bis ein Wert ankommt, dann geht es weiter und dann endet die Goroutine

		fmt.Println("Received", amount, "$!")
	}()

	// 2. Send an Channgel money
	money <- 200
}
