package main

import "fmt"

func f() int {
	i := 5
	defer func() {
		i++
	}()
	return i
}

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return r
}

func f4() {
	defer func() {
		fmt.Println("444")

	}()

	defer func() {
		fmt.Println("555")
	}()

	println("dddddd")
}

func main() {
	//println(f())
	//println(f1())
	//println(f2())
	//println(f3())

	f4()
}
