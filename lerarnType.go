package main

import "fmt"

type myFunc func(int) int

type myStruct struct {
}

func (f myFunc) sum(a int, b int) int {
	res := a + b
	return res
}

func sum100(a int) int {
	return a * 100
}

func sum10(a int) int {
	return a * 10
}

func handleSum(handler myFunc, a int, b int) int {
	res := handler.sum(a, b)
	return handler(res)
}

func main() {

	newFunc1 := myFunc(sum10)
	res1 := handleSum(newFunc1, 1, 1)
	res2 := handleSum(sum100, 1, 1)

	fmt.Println(res1)
	fmt.Println(res2)

}
