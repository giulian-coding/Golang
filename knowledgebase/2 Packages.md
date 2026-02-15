# Packages

Every Go program is made up of packages. The Program starts in the package "main"

## Was ist ein Package?

Ein Package ist eine Sammlung von Go-Dateien im selben Verzeichnis, die gemeinsam kompiliert werden. Jedes Go-File beginnt mit einer `package`-Deklaration.

```go
package math
```

Das Hauptprogramm startet immer im Package `main` und benötigt eine Funktion `func main() { ... }` als Einstiegspunkt.

## Import von Packages

Um Funktionen aus anderen Packages zu verwenden, nutzt man das `import`-Statement:

```go
import (
	"fmt"      // Standardbibliothek
	"meinprojekt/mypkg" // eigenes Package
)
```

Danach können exportierte Funktionen (mit Großbuchstaben beginnend) genutzt werden:

```go
fmt.Println("Hallo Welt")
mypkg.MyFunction()
```

## Eigenes Package erstellen

1. Lege ein neues Verzeichnis an, z.B. `mypkg`.
2. Erstelle darin eine Datei, z.B. `util.go`:

```go
package mypkg

func MyFunction() {
	// ...
}
```

3. Importiere das Package im Hauptprogramm (siehe oben).

## Sichtbarkeit von Namen

- Namen, die mit einem Großbuchstaben beginnen, sind **exportiert** (öffentlich).
- Namen mit Kleinbuchstaben sind **nicht exportiert** (privat für das Package).

## Beispiel: main.go und eigenes Package

**mypkg/util.go:**

```go
package mypkg

func Add(a, b int) int {
	return a + b
}
```

**main.go:**

```go
package main

import (
	"fmt"
	"meinprojekt/mypkg"
)

func main() {
	result := mypkg.Add(2, 3)
	fmt.Println(result)
}
```

## Weitere Hinweise

- Der Importpfad für eigene Packages richtet sich nach dem Modulnamen in der `go.mod`.
- Mit `go doc` können Informationen zu Packages angezeigt werden.
