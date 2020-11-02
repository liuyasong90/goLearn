package main

import "fmt"

type testStruct struct {
}

func (testStruct) testFunc01() {
	fmt.Println("7777")
}

func main() {

	a := testStruct{}
	a.testFunc01()
	fmt.Print("4")
}
