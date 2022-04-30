package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
)

var base58Alphabets = []byte("zpxhncf39wBUDNEGHJKLM4PQRST7VWXYZ2badeCg65jkm8oFqi1tuvAsyr")

func main() {
	// testReverse()

	str := []byte{28, 3, 71, 104, 250, 55, 8, 198, 176, 160, 235, 96, 191, 95, 59, 58, 177, 139, 97, 145, 248, 191, 115, 198, 12, 135, 132, 77, 171, 70, 151, 87, 204, 105}
	_, res := Base58Encode(str)
	fmt.Println("res=", res)
	resByte, resStr := Base58Decode([]byte(res))
	fmt.Println("resByte=", resByte)
	fmt.Println("resStr=", resStr)

}

// Base58Encode 编码
func Base58Encode(input []byte) ([]byte, string) {
	x := big.NewInt(0).SetBytes(input)
	fmt.Println("x=", x)
	base := big.NewInt(58)
	zero := big.NewInt(0)
	mod := &big.Int{}
	var result []byte
	// 被除数/除数=商……余数
	fmt.Println("开始循环-------")
	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		fmt.Println("mod=", mod)
		fmt.Println("x=", x)
		result = append(result, base58Alphabets[mod.Int64()])
		fmt.Println("一次循环结束-------")
	}
	ReverseBytes(result)
	return result, string(result)
}

// Base58Decode 解码
func Base58Decode(input []byte) ([]byte, string) {
	result := big.NewInt(0)
	for _, b := range input {
		charIndex := bytes.IndexByte(base58Alphabets, b)
		result.Mul(result, big.NewInt(58))
		result.Add(result, big.NewInt(int64(charIndex)))
	}
	decoded := result.Bytes()
	if input[0] == base58Alphabets[0] {
		decoded = append([]byte{0x00}, decoded...)
	}
	return decoded, string(decoded)
}

func testReverse() {
	str := "12345678"
	// data := []byte(str)
	data, _ := hex.DecodeString(str)
	fmt.Println(data)
	ReverseBytes(data)
	fmt.Println(fmt.Sprintf("%v", data))
}

// ReverseBytes 翻转字节
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
