package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.7
	y := 5
	fmt.Println(x / float64(y))

	nota := 7.5
	notaf := int(nota)
	fmt.Println(notaf)
	//cuidado isso não faz a msm coisa  do int pro str
	fmt.Println("teste  " + string(97))

	// inteiro para string
	fmt.Println("teste" + strconv.Itoa(123))
	// string para inteiro
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("verdadeiro")
	}

}
