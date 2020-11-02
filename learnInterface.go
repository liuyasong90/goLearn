package main

import "fmt"

type inter01 interface {
	Method01(int)
	Method02(string, int)
}

type struct02 struct {
	inter inter01
}

type TestInter03 struct {
}

func (t TestInter03) Method01(a int) {

	fmt.Println(a)
}

func (t TestInter03) Method02(s string, a int) {

	fmt.Println(a)
	fmt.Println(s)

}

func fanfa1(a inter01, i int) {
	a.Method01(i)
	a.Method02("333", i)
}

func main() {

	t3 := TestInter03{}

	fanfa1(t3, 1)

}
