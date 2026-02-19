package main

import (
	"fmt"
	"time"
)

func main() {
	for a := 1; a <= 10; a++ {
		fmt.Println(a)
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}
	fmt.Println("\nMisturando...")
	for a := 1; a <= 10; a++ {
		if a%2 == 0 {
			fmt.Println("é numero par")
		} else {
			fmt.Println("é numero impar")
		}
	}
	for {
		fmt.Println("para sempre...") //este é um ex de laço infinito de for
		time.Sleep(time.Second * 20)  // use time.Sleep, not fmt.Sleep
	}

}
