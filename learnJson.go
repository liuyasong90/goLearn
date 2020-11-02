package main

import (
	"encoding/json"
	"fmt"
)

func test() {
	m := map[string]interface{}{"a": 1, "b": 2, "c": 3}
	fmt.Println(m)

	data, _ := json.Marshal(m)
	fmt.Println(string(data))
	fmt.Println(data)

	m1 := make(map[string]interface{})
	_ = json.Unmarshal(data, &m1)
	fmt.Println(m1)

}

func main() {

	test()

}
