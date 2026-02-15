package main

//

import "fmt"

// global variable
var aPackageVar string = "a package variable"

const greetingPart2 = "Gopher"

// shadowed variable
var shadow1 string = "a shadowed string"

func main() {
	// Deklaration und Zuweisung in langer Form
	var s1 string = "a string"

	// Deklaration und Zuweisung in kurzer Form
	s2 := "another string"

	// Deklaration un Zuweisung trennen
	var s3 string
	s3 = "a third string"

	var s0 string
	var i0 int
	var f0 float64
	var b0 bool

	fmt.Println(s0, i0, f0, b0)
	fmt.Println(s1, s2, s3, aPackageVar)

	// Konstanten
	const greetingPart1 = "Hello"

	fmt.Println(greetingPart1, greetingPart2)

	fmt.Println(shadow1)

	// shadowing
	shadow1 = "a new shadowed string"
	fmt.Println(shadow1)

	// Arrays
	names := [2]string{"Alice", "Bob"}
	fmt.Println(names)
	names[1] = "Charlie"
	fmt.Println(names)

	// Slices mit make
	numbers := make([]string, 3)
	numbers[0] = "one"
	numbers[1] = "two"
	numbers[2] = "three"
	fmt.Println(numbers)

	// Slices mit Literalen
	moreNumbers := []int{4, 5, 6}
	fmt.Println(moreNumbers)

	//Neue Werte zu einem Slice hinzufügen
	moreNumbers = append(moreNumbers, 7, 8)
	fmt.Println(moreNumbers)

	// Maps
	// Deklaration und Zuweisung mit make
	age := make(map[string]int)
	age["Alice"] = 30
	age["Bob"] = 25
	fmt.Println(age)

	// Deklaration und Zuweisung mit Literalen
	capitals := map[string]string{
		"Germany": "Berlin",
		"France":  "Paris",
	}
	fmt.Println(capitals)

	// Zugriff auf Map-Werte
	fmt.Println("Capital of Germany:", capitals["Germany"])

	// Überprüfen, ob ein Schlüssel existiert
	capital, exists := capitals["Spain"]
	if exists {
		fmt.Println("Capital of Spain:", capital)
	} else {
		fmt.Println("Capital of Spain not found")
	}

	// Structs
	// Structs sind benutzerdefinierte Datentypen, die mehrere Felder enthalten können
	type Person struct {
		salary float64 // unexported field
		Name   string  // exported field
		Age    int
	}

	type Employee struct {
		Person   // embedding
		Position string
	}

	// Struct-Instanz mit Literalen erstellen
	alice := Person{Name: "Alice", Age: 30, salary: 50000}
	fmt.Println(alice)

	peter := Employee{Position: "Manager"}
	fmt.Println(peter)
	peter.Name = "Peter"
	peter.Age = 40
	peter.salary = 80000
	fmt.Println(peter)

	// Zugriff auf Struct-Felder
	fmt.Println("Alice's name:", alice.Name)
	fmt.Println("Alice's age:", alice.Age)

	// Leere Struct Instanz
	var emptyPerson Person
	fmt.Println(emptyPerson)

	// Pointer
	var p *Person
	p = &alice
	fmt.Println("Pointer to Alice:", p)
	fmt.Println("Alice's name via pointer:", p.Name)

	// Pointer auf Int
	var num int = 42
	var numPtr *int = &num
	fmt.Println("Pointer to num:", numPtr)
	fmt.Println("Value of num via pointer:", *numPtr)

	// Pointer mit * auflösen und Wert zuweisen
	*numPtr = 100
	fmt.Println("Updated num:", num)

	(*p).Age = 31 // Alternative Schreibweise
	fmt.Println("Updated Alice's age:", alice.Age)

	// Leerer Wert nil
	var t *Person
	fmt.Println("Nil pointer:", t)

	var s []float64
	fmt.Println("Nil slice:", s)

	var m map[string]int
	fmt.Println("Nil map:", m)

	fmt.Println("Nil pointer check:", t == nil)

	// Error
	err := fmt.Errorf("user %v not found", "Frank")
	fmt.Println("Error err:", err)

	
}
