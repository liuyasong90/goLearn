package main

import "strconv"
import "fmt"

func main() {
	str := "11"
	S, _ := strconv.Atoi(str)
	fmt.Println([]byte(strconv.Itoa(S)))

}
