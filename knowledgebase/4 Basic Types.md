# Go Datentypen (Basics)

Go besitzt mehrere eingebaute Datentypen, die für unterschiedliche Zwecke verwendet werden.

## C-Syntax (link-to-left Lesen)

In C werden Typen von rechts nach links gelesen. Beispiel:

```c
int (*fp)(int (*)(int, int), int);
```

Lesart:

> `fp` ist ein Zeiger auf eine Funktion, die einen Zeiger auf eine Funktion (die zwei `int`-Parameter und einen `int`-Rückgabewert hat) und einen `int`-Parameter nimmt und einen `int`-Rückgabewert hat.

Vorteile:

- mächtig und flexibel
- ermöglicht komplexe Typen

Nachteile:

- schwer erkennbar, **was eigentlich deklariert wird**
- unklar, wo der Name eingefügt werden muss
- Parsing komplizierter (→ Klammern bei Casts nötig)

Go vermeidet diese Probleme durch eine klarere Syntax.

## Go: Basistypen

Go bietet folgende eingebaute Basistypen:

| Typ          | Beschreibung                  | Beispielwert  |
| ------------ | ----------------------------- | ------------- |
| `bool`       | Wahrheitswert (true/false)    | `true`        |
| `string`     | Zeichenkette                  | `"Hallo"`     |
| `int`        | Ganzzahl (plattformabhängig)  | `42`          |
| `int8`       | 8-Bit-Ganzzahl                | `-128`        |
| `int16`      | 16-Bit-Ganzzahl               | `32767`       |
| `int32`      | 32-Bit-Ganzzahl               | `-2147483648` |
| `int64`      | 64-Bit-Ganzzahl               | `1234567890`  |
| `uint8`      | 8-Bit-unsigned (alias `byte`) | `255`         |
| `float32`    | 32-Bit-Gleitkommazahl         | `3.14`        |
| `float64`    | 64-Bit-Gleitkommazahl         | `2.71828`     |
| `complex64`  | Komplexe Zahl (float32)       | `1+2i`        |
| `complex128` | Komplexe Zahl (float64)       | `2-3i`        |

Weitere Typen: Arrays, Slices, Structs, Maps, Channels, Interfaces, Funktionen, Zeiger.

## Variablendeklaration in Go

```go
var x int = 42      // explizite Typangabe
var y = "Hallo"     // Typinferenz (string)
z := true           // Kurzschreibweise, Typ: bool
```

## Beispiel: Funktionen als Typen

```go
var f func(int, int) int
```

`f` ist eine Variable, die eine Funktion speichert, die zwei `int`-Parameter nimmt und einen `int` zurückgibt.

## Vorteile der Go-Syntax

- Lesbarkeit: Name steht immer links, Typ immer rechts
- Keine Klammern für Zeiger/Funktionen nötig
- Einfaches Parsing, weniger Fehlerquellen

```go

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
```

Output:

```
Type: bool Value: false
Type: uint64 Value: 18446744073709551615
Type: complex128 Value: (2+3i)
```

## Zero Values

In Go haben Variablen immer einen definierten Wert, auch wenn sie nicht explizit initialisiert wurden. Dieser Wert wird als "Zero Value" bezeichnet.

| Typ                                            | Zero Value                |
| ---------------------------------------------- | ------------------------- |
| `bool`                                         | `false`                   |
| `string`                                       | `""` (leere Zeichenkette) |
| `int`, `uint`                                  | `0`                       |
| `float`                                        | `0.0`                     |
| `complex`                                      | `0 + 0i`                  |
| Zeiger                                         | `nil`                     |
| Slices, Maps, Channels, Interfaces, Funktionen | `nil`                     |

## Type conversions

Go unterstützt explizite Typumwandlungen, aber keine impliziten. Beispiel:

```go
var i int = 42
var f float64 = float64(i) // explizite Umwandlung von int zu float64
var u uint = uint(f) // explizite Umwandlung von float64 zu uint (möglicher Datenverlust)

i := 42
f := float64(i)
u := uint(f)
```

Go erlaubt keine impliziten Umwandlungen, um unerwartete Fehler zu vermeiden. Alle Typumwandlungen müssen explizit angegeben werden.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
```

Output:

```
3 4 5
```

## Type inference

Go unterstützt Typinferenz, was bedeutet, dass der Compiler den Typ einer Variable basierend auf dem zugewiesenen Wert (rechte Seite der Deklaration) ableiten kann. Beispiel:

```go
package main

import "fmt"

func main() {
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
}
```

Output:

```
v is of type int
```

In diesem Beispiel wird der Typ von `v` automatisch als `int` erkannt, da der zugewiesene Wert `42` ein Ganzzahl-Literal ist.
Die Typinferenz erleichtert die Deklaration von Variablen, da der Entwickler nicht immer den Typ explizit angeben muss. Es ist jedoch wichtig zu beachten, dass die Typinferenz nur funktioniert, wenn der Compiler genügend Informationen hat, um den Typ eindeutig zu bestimmen.

## Constants

In Go können Konstanten mit dem Schlüsselwort `const` definiert werden. Konstanten sind unveränderliche Werte, die zur Kompilierzeit festgelegt werden. Beispiel:

```go
package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
```

Konstanten können auch unbenannte Werte haben, die automatisch inkrementiert werden. Beispiel:

```go
package main

import "fmt"

const (
    _  = iota             // ignore first value by assigning to blank identifier
    a  = iota             // a == 1
    b  = iota             // b == 2
    c  = iota             // c == 3
)
func main() {
    fmt.Println(a, b, c) // Output: 1 2 3
}
```

In diesem Beispiel wird `iota` verwendet, um automatisch inkrementierende Werte für die Konstanten `a`, `b` und `c` zu generieren. Die erste Zeile `_ = iota` ignoriert den Wert `0`, sodass `a` mit `1` beginnt.

## Numeric Constants

Go unterstützt auch numerische Konstanten, die unendlich genau sein können. Diese Konstanten können als Ganzzahlen, Gleitkommazahlen oder komplexe Zahlen dargestellt werden. Beispiel:

```go
package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
```

Output:

```
21
0.2
1.2676506002282294e+30
```
