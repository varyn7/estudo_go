package main

import "fmt"

func main() {
	fmt.Print("Hello, World!")
	fmt.Print("ate mais !")

	fmt.Println(" how are you ?")
	x := 3.14
	xs := fmt.Sprintf("%g", x)
	fmt.Println(xs, "esse é o valor de x")

	a := 2007
	b := "pera"
	c := false
	d := 1.999999
	fmt.Printf("\n%d %f %.1f %t %s", a, b, c, d)

}
