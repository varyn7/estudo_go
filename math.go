package main

import (
	"fmt"
	"math"
)

func main() {
	pi := 3.1415926
	r := 2.0
	a := pi * math.Pow(r, 2)
	fmt.Println(a)
}
