package main

import "fmt"

type Handle01 interface {
	Fn(a int, b int)
}

type HandleFunc01 func(int) int

func (f HandleFunc01) Fn(a int, b int) {
	fmt.Println("begin ....")
	f(a)
	fmt.Println("end......")
}

type Home struct {
	handle HandleFunc01
}

func Func01Tpye01(a int) int {
	fmt.Println(a, "Type01")
	return 1
}

func Func01Tpye02(a int) int {
	fmt.Println(a, "Type02")

	return 1
}

func main() {
	sturct01 := Home{handle: Func01Tpye01}

	sturct01.handle.Fn(1, 2)

	sturct02 := Home{handle: Func01Tpye02}
	sturct02.handle.Fn(1, 2)

}
