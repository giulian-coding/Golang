package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	// Pi and Pizza are exportet --> pi and pizza aren't

	fmt.Println(add(42, 13))

	a, b := swap("Hello", "World")
	fmt.Println(a, b)

	fmt.Println(split(17))

	var i int
	fmt.Println(i, c, python, java)

	// Function Scope - lokal -- shadowing
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	//Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	//Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
	var g, j int = 1, 2
	k := 3
	c, python, java = false, true, "nod!"

	fmt.Println(g, j, k, c, python, java)
}

// private function
func add(x int, y int) int {
	return x + y
}

// public function
func Add(x, y int) int { // shortened typing
	return x + y
}

func swap(x, y string) (string, string) {
	return x, y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// List of Variables
// Package Scope - global
var c, python, java bool

var i, j int = 1, 2
