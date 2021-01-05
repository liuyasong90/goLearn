package main

import "fmt"

func assertTest(a interface{}) {
	println(a.(int))
}

func main() {
	assertTest(1)
	fmt.Println("test merge")
}
