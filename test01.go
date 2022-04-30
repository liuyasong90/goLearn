package main

import (
	"fmt"
	"strconv"
)

type testStruct struct {
}

func (testStruct) testFunc01() {
	fmt.Println("7777")
}

// 0203174510E3FD8265D88758F8D6A762B0F2507690143DC0F2456102795E9E296C
func main() {
	//publicKey := "n9JZiFFQkQdZzD161cAonZcJMNxTpQwoKvJPt9m853EXEKqLCJy1"
	//byteData := []byte(publicKey)
	//hexStringData := hex.EncodeToString(byteData)
	//print(hexStringData)
	//fmt.Printf("\n%X", publicKey)
	//a := int(100)
	//print("dddddddddddddddddddddddd\n")
	//print("ssssssssssssss     " + strconv.Itoa(a) + "  weeee")

	str1 := "10001"

	nt, _ := strconv.ParseFloat(str1, 64)
	st := fmt.Sprintf("%.6f", nt/1000000)
	print(st + "ZXC")

}
