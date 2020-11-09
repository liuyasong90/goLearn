package main

import (
	"encoding/json"
	"fmt"
)

func test1() {
	m := map[string]interface{}{"name": "taoge", "age": 30, "addr": "China"}
	fmt.Println(m)

	data, _ := json.Marshal(m)
	fmt.Println(string(data))

	m1 := make(map[string]interface{})
	_ = json.Unmarshal(data, &m1)
	fmt.Println(m1)
}

func main() {
	test1()
	fmt.Println("dddd")
}
