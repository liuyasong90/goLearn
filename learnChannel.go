package main

import (
	"strings"
)

//func main() {
//	defer func() {
//		fmt.Println(err)  }()
//	defer func() {  fmt.Println("a") }()
//	defer func() { fmt.Println("b")
//	}()
//	panic("c")
//	}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	firstStr := strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], firstStr) != 0 {
			firstStr = firstStr[:len(firstStr)-1]
		}
	}
	if len(firstStr) > 0 {
		return firstStr
	}
	return ""
}

//func commonPrefix(a, b string) string {
//
//	commentStr:=a
//	for strings.Index(b, commentStr) != 0 {
//		commentStr = commentStr[:len(commentStr) - 1]
//	}
//
//	if len(commentStr)>0{
//		return commentStr
//	}
//	return ""
//}
//
//func main() {
//	//s := []string{"liu","li"}
//	fmt.Println(commonPrefix("li","l"))
//
//}

var urls = []string{"http://baidu.com", "http://sogo.com", "http://bing.com", ""}
var result []string

func main() {

}
