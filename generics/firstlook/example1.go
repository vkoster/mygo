package main

import "fmt"

/*
Ein erstes Beispiel: eine Funktion zum Addieren unterschiedlicher nummerischer Typen.
*/

// früher brauchte man eine Funktion pro Type:
func addI(a, b int) int {
	return a + b
}
func addF(a, b float64) float64 {
	return a + b
}

// dafür kann man eine generische Function schreiben
func addT[T int | float64](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(addI(1, 2))
	fmt.Println(addF(1.5, 2.5))

	fmt.Println(addT(1, 2))
	fmt.Println(addT(1.5, 2.5))

	// das geht immer noch nicht!
	// Aber jetzt kommt der Fehler erst zur Laufzeit - der Compiler kann das nicht mehr checken...
	// fmt.Println(addI(1, 2.5))
}
