package main

import (
	"fmt"
)

func testb(x int) {

	var a [10]int
	a[x] = 10
}

func test08() {
	fmt.Println("cccc")
}
func main() {

	testb(11)
	test08()

}
