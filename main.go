package main

//

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Congressman struct {
	Name           string
	AccountBalance float64
}

func (c Congressman) greet() {
	fmt.Println("Hello", c.Name)
}

type Enemy struct{}

func (e Enemy) greet() {
	fmt.Println("Go to Hell!")
}

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

	// If/Else
	president := "Walker"
	if president != "Underwood" {
		fmt.Println("Underwood is not president")
	}

	var proVoteCount int
	if len(os.Args) < 2 {
		proVoteCount = 300 // Standardwert
	} else {
		arg := os.Args[1]
		num2, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Ungültiger Parameter, bitte eine Zahl eingeben.")
			return
		}
		proVoteCount = num2
	}
	if proVoteCount < 300 {
		fmt.Println("Not enough")
	} else if proVoteCount == 301 {
		fmt.Println("Draw")
	} else {
		fmt.Println("We've got the votes")
	}

	// Switch
	// mit Expressions
	proVoteCount2 := 301
	switch {
	case proVoteCount2 < 300:
		fmt.Println("Not enough")
	case proVoteCount2 == 301:
		fmt.Println("Draw")
	default:
		fmt.Println("We've got the votes")
	}

	// über Typ
	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("Monday")
	default:
		fmt.Println("No monday")
	}

	// Schleifen
	// Klassisch
	for i2 := 1; i2 < 7; i2++ {
		fmt.Println(i2)
	}

	// While
	j := 1
	for j < 7 {
		fmt.Println("Season", j)
		j++
	}
	// Endlos
	/*
		k := 1
		for {
			fmt.Println("Season", k)
			k++
		}
	*/

	// For-Each mit range
	supportingCharacters := []string{"Freddy Hayes", "Donald Blythe"}

	for index, character := range supportingCharacters {
		fmt.Println("Character", character, "index", index)
	}

	for dogAge := 1; dogAge < 11; dogAge++ {
		humanAge := dogAge * 7
		fmt.Printf("Dog: %3d Human: %5d\n", dogAge, humanAge)
	}

	hassoAge := 3
	fmt.Println("Human Age of Hasso", humanAge(hassoAge))

	human1, human2 := humanAges(3, 5)
	fmt.Println("Human ages:", human1, ",", human2)

	c := Congressman{Name: "Peter", AccountBalance: 8000.0}

	c.swearOathOfOffice()
	err3 := c.bribe(5000.0)
	if err3 != nil {
		fmt.Println("Could not bribe congressman:", err3)
	}

	var i3 int = humanAge(3)
	log.Println(i3)

	var j3 float64 = humanAge(4.9)
	log.Println(j3)

	// anonyme Funktion -> zb für HTTP Handler
	func() {
		fmt.Println("Hello LinkedIn")
	}()

	helloDefer()

	e := Enemy{}

	passBy(c, e)

	fmt.Println(c)

	// Das leere Interface
	var i4 interface{} = "Hello!"

	// Type Assertion ohne Prüfung
	var s4 string = i4.(string)
	fmt.Println(s4)

	// Type Assertion mit Prüfung
	s2, ok := i4.(string)
	fmt.Println(s2, "is a string?", ok)

	// Falsche Type Assertion
	// var f float64 = i4.(float64)
	// fmt.Println(f)
	// panic: interface conversion: interface {} is string, not float64
}

// Funktionen
// Einfach
// func humanAge(dogAge int) int {
// 	return dogAge * 7
// }

// Oder so
// func humanAge(dogAge int) int {
// 	humanAge := dogAge * 7
// 	return humanAge
// }

// Benannter Return Wert, "naked return"
// func humanAge(dogAge int) (humanAge int) {
// 	humanAge = dogAge * 7
// 	return
// }

// kleiner Buchstabe --> private Funktion
// func humanAge

//grosser Buchstabe --> öffentliche Funktion
// func HumanAge

func humanAges(dog1, dog2 int) (int, int) {
	return humanAge(dog1), humanAge(dog2)
}

func (c Congressman) swearOathOfOffice() {
	fmt.Println("I", c.Name, "swear to serve the US!")
}

func (c Congressman) bribe(amount float64) error { // error ist hier einfach ein return value
	if c.Name != "Peter" {
		return fmt.Errorf("%v is not corrupt", c.Name)
	}
	c.AccountBalance += amount
	fmt.Println(c.Name, "has", c.AccountBalance)
	return nil
}

// humagAge Funktion war nur mit int kompatibel, aber evtl will man mit float und int
// generische Funktion
func humanAge[T int | float64](dogAge T) (humanAge T) {
	humanAge = dogAge * 7
	return
}

// mit defer keyword kann ausführung von Code verzögert werden -> bis zum Ende der Funktion
// defer wird immer ausgeführt, auch wenn Fehler innerhalb der Funktion auftreten
func helloDefer() {
	defer fmt.Print("LinkedIn\n")

	// var s3 = []string{}
	// fmt.Println("Ouch ", s3[42])
	fmt.Print("Hello ")
}

// Interfaces
// Interface definieren

type Greeter interface {
	greet()
}

// Interface nutzen
func passBy(g1, g2 Greeter) {
	g1.greet()
	g2.greet()
}

func (c Congressman) String() string {
	return "Hello! " + c.Name
}
