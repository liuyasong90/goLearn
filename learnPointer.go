package main

import "fmt"

// 声明整形切片指针类型，表示该类型变量里存储的是一个切片的首地址
type arrPoint *[]int

func test01() {
	// 声明并初始化数组，大小为6个元素整形数组，前5个数字是0，最后一个2
	var a = []int{5: 2}
	var a1 = []int{5: 3}
	// 数组指针，变量的内容是数组的地址

	var b arrPoint = &a
	var b1 arrPoint = &a1
	// 切片数组，大小为2元素类型为arrPiont(指针)，切片里存储元素的类型是arrPiont指针
	c := []arrPoint{b, b1}
	// 也可以这样写
	// c := [2]arrPoint(&a, &a1)
	for _, v := range c {
		fmt.Println(*v)
	}
}

func test03() {

	a := []int{1, 2, 3}

	var b *[]int
	b = &a
	a[0] = 4

	for i := 0; i < 10000; i++ {
		a = append(a, i)
	}
	(*b)[2] = 8
	for i, value := range a {
		println(value)

		if i == 10 {
			break
		}
	}

	println("4hjhj")
	println((*b)[2])

	//println(a)
	//println(*b)

}

//指针切片
func test02() {
	a := 1
	b := 2
	c := 4
	//
	//k:=make([]*int,2)
	//
	//k=append(k, &a)
	//k=append(k,&b)

	k := []*int{&a, &b}
	k = append(k, &c)
	*k[2] = 19

	println(k)
	for _, value := range k {
		*value = *value + 1
		println(*value)
	}
}
func main() {
	//test02()

	test03()

}
