package main

import "fmt"

func main() {
	// Estrutura homogenea e estatica fixa: exemplo de array
	var notas [3]float64
	notas[0] = 10.0
	notas[1] = 2.5
	notas[2] = 3.5
	fmt.Println(notas)
	total := 0.0
	for _, v := range notas {
		total += v
	}
	media := total / float64(len(notas))
	fmt.Println("total:", total)
	fmt.Println("media:", media)
}
