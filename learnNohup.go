package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("333")
		time.Sleep(1 * time.Second)
	}
}
