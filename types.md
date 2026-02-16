Erklärung Golang Types

Kurz gesagt: **Methoden gibt’s in Go für jeden *benannten* Typ – nicht nur für `struct`s.**  
Aber: Du kannst **nicht** direkt auf eingebaute Typen wie `int` Methoden definieren.

### Warum geht `func (i int) Greet()` nicht?

Weil `int` ein **nicht-lokaler**, eingebaulter Typ ist. In Go gilt:

*   Eine **Methode** ist jede Funktion mit **Receiver** (z. B. `(p Person)` oder `(*Person)` oder `(m MyInt)`).
*   **Receiver-Typen müssen benannte Typen sein, die im *gleichen Paket* definiert wurden.**
*   Auf **eingebauten Typen** (wie `int`, `string`, `map[...]...`) oder **Fremdtypen aus anderen Paketen** darfst du **keine** Methoden definieren.

Daher ist das hier **illegal**:

```go
func (i int) Greet() {} // ❌ compile error: cannot define new methods on non-local type int
```

### So geht’s richtig: eigenen Typ definieren

Wenn du eine Methode „auf int“ willst, mach erst einen **benannten eigenen Typ** mit `int` als *Underlying Type*:

```go
package main

import "fmt"

type MyInt int // eigener benannter Typ mit underlying type int

func (i MyInt) Greet() {
    fmt.Println("Hallo aus MyInt:", i)
}

func main() {
    var x MyInt = 42
    x.Greet() // -> Hallo aus MyInt: 42
}
```

Das funktioniert, weil `MyInt` ein **lokal definierter, benannter Typ** ist.

### Achtung: Typalias vs. neuer Typ

*   **Neuer Typ**:
    ```go
    type MyInt int // ✅ eigener Typ -> Methoden möglich
    ```
*   **Alias**:
    ```go
    type MyInt = int // ❌ nur ein Alias für int -> Methoden NICHT möglich
    ```

### Gilt nicht nur für `struct`

Du kannst Methoden auf beliebigen benannten Typen definieren, z. B.:

```go
type ID string
type Counter int
type StringList []string
type KV map[string]int

func (id ID) Valid() bool      { return len(id) > 0 }
func (c *Counter) Inc()        { *c++ }
func (s StringList) First() string { if len(s) > 0 { return s[0] }; return "" }
func (m KV) Keys() []string {
    keys := make([]string, 0, len(m))
    for k := range m { keys = append(keys, k) }
    return keys
}
```

> Wichtig: `StringList` und `KV` sind **benannte** Typen, daher sind Methoden erlaubt – obwohl die underlying types `[]string` bzw. `map[string]int` sind.

### Value- vs. Pointer-Receiver (kurz & praxisnah)

*   **Value-Receiver** `(t T)`: Methode arbeitet mit einer **Kopie**; gut für kleine, unveränderliche Werte.
*   **Pointer-Receiver** `(t *T)`: Methode kann **ändern** und vermeidet Kopien; Standard bei größeren Structs oder wenn die Methode den Zustand ändern soll.

Go ruft automatisch passend auf:

```go
p := Person{Name: "Giulian"} 
p.SetName("Gigi")   // funktioniert, auch wenn SetName einen *Person-Receiver hat
(&p).SetName("Gigi") // äquivalent
```

### Method Set & Interfaces (ein Satz für Fortgeschrittene)

*   Der **Method Set** eines Typs bestimmt, welche Interfaces er erfüllt.
*   `T` hat nur Methoden mit **Value-Receiver**.
*   `*T` hat Methoden mit **Value- und Pointer-Receiver**.
*   Daher erfüllt `*T` häufiger Interfaces als `T`.

***

## Fazit

*   **Methoden** = jede Funktion mit **Receiver**.
*   **Erlaubt** auf: **benannten Typen**, die du **im gleichen Paket** definierst (egal ob `struct`, `int`-basiert, `map`-basiert, etc.).
*   **Nicht erlaubt** auf: **eingebauten Typen direkt** (`int`, `string`, …) oder auf **Aliasen** davon.

Wenn du magst, zeige ich dir ein Mini-Projekt, wo wir einem `type MyInt int` ein paar nützliche Methoden geben (z. B. `IsEven`, `Clamp`, `Format`) und das mit Interfaces koppeln. Soll ich das kurz aufsetzen?
