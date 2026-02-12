package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("strings  ", "pão" == "pão")

	//operador relacional de diferença

	fmt.Println("!=", 3 != 2)
	//operador de comparação de tamanho

	fmt.Println("<", 3 < 2)
	//operador de comparação de tamanho
	fmt.Println(">", 3 > 2)
	//operador de menor ou igual
	fmt.Println("<=", 3 <= 2)
	//operador: maior ou igual
	fmt.Println(">=", 3 >= 2)
	//comparativo de tempo
	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)
	fmt.Println("datas :", d1 == d2)
	//ou
	fmt.Println("datas :", d1.Equal(d2))

}
