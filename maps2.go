package main

import "fmt"

func main() {
	funcsesalarios := map[string]float64{
		"fulano":   123456,
		"ciclano":  124450,
		"beltrano": 148800,
	}
	funcsesalarios["rafael menino"] = 12000
	delete(funcsesalarios, "inexistente")

	for nome, salario := range funcsesalarios {
		fmt.Println(nome, salario)
	}
}
