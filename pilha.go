package main

import "runtime/debug"

func stackF3() {
	debug.PrintStack()
}
func stackF2() {
	stackF3()
}
func stackF1() {
	stackF2()
}
func main() {
	stackF1()
}
