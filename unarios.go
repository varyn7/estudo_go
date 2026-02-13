package main

import "fmt"

func main() {
	x := 1
	y := 2

	//em golang só é permitido postfix
	x++
	fmt.Println(x) //x=1 ou x=x+1
	y--
	fmt.Println(y) //y=1 ou y=-1
	fmt.Println(x == y)
}
