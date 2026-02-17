package main

import "fmt"

func main() {
	tests := []float64{10, 9.5, 9, 8.9, 8, 7.5, 5, 4, 3, 2.9, 0}
	for _, n := range tests {
		fmt.Printf("n=%.1f -> %s\n", n, Notaparateste(n))
	}
}
func Notaparateste(n float64) string {
	if n >= 9 && n <= 10 {
		return "a"
	} else if n >= 8 && n < 9 {
		return "b"
	} else if n >= 5 && n < 8 {
		return "c"
	} else if n >= 3 && n < 5 {
		return "d"
	} else {
		return "e"
	}
}
