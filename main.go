package main

import (
	"fmt"
)

func main() {
	s := "BrainWu"
	if v, ok := interface{}(s).(string); ok {
		fmt.Println(v)
	}
}
