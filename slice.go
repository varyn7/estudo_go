package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [4]int{1, 2, 3, 4} //array
	s1 := []int{1, 2, 3}     //slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))
	a2 := [6]int{1, 2, 3, 4, 5, 6}
	s2 := a2[1:3]
	fmt.Println(a2, s2)
	s3 := a2[:2] //novo slice mas aponta para o mesmo array
	fmt.Println(a2, s3)
	//voce pode imaginar slice como tamano e um ponteiro para um elemento de um array
	s4 := s2[:1]
	fmt.Println(a2, s4)
}
