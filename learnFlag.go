package main

import (
	"flag"
	"fmt"
)

var (
	name    string
	age     int
	address *string
	id      *int
)

func init() {
	flag.StringVar(&name,
		"name",
		"匿名",
		"您的姓名")

	flag.IntVar(&age,
		"age",
		-1,
		"您的年龄")

	address = flag.String("address",
		"未知",
		"您的住址")
	id = flag.Int("id", -1, "身份ID")
}

func main() {
	flag.Parse()

	fmt.Printf("%s您好, 您的年龄:%d, 您的住址:%s, 您的ID:%d\n\n", name, age, *address, *id)

	fmt.Println("--遍历有输入的参数（开始）---")
	flag.Visit(func(f *flag.Flag) {
		fmt.Printf("参数名[%s], 参数值[%s], 默认值[%s], 描述信息[%s]\n\n", f.Name, f.Value, f.DefValue, f.Usage)
	})
	fmt.Println("--遍历有输入的参数（结束）--\n")

	fmt.Println("--遍历所有参数（开始）--\n")
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("参数名[%s], 参数值[%s], 默认值[%s], 描述信息[%s]\n\n", f.Name, f.Value, f.DefValue, f.Usage)

	})
	fmt.Println("--遍历所有参数（结束）--\n")

}
