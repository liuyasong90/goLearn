package main

import "fmt"

type TestInter02 struct {
}

func (t *TestInter02) Method01(a int) {

	fmt.Println(a)
}

func (t *TestInter02) Mthod02(s string, a int) {

	fmt.Println(a)
	fmt.Println(s)
}
