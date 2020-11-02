package main

import (
	"fmt"
	"github.com/goinggo/mapstructure"
)

type Entity struct {
	Num int
	S   string
	T   map[string]string
}

func main() {
	var te Entity
	m := make(map[string]interface{})
	m["Num"] = 1
	m["S"] = "test"
	m["T"] = map[string]string{"1": "1", "2": "2"}

	err := mapstructure.Decode(m, &te)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(te)
	fmt.Print(te.Num, " ", te.S, " ", te.T)
}
