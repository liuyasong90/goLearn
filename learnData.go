package main

import "fmt"

func slice() {
	var slice01 []int

	slice01 = append(slice01, 1)

	_ = slice01

	fmt.Print(slice01[0])
}

func main() {
	slice()

}
