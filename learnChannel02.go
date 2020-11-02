package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	stop := make(chan int)
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go func() {
		for {
			select {
			case <-stop:
				time.Sleep(2 * time.Second)
				fmt.Println("协程结束......")
				waitGroup.Done()
				return
			default:
				fmt.Println("等待中......")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	stop <- 1
	fmt.Println("可以了结束协程......")

	waitGroup.Wait()

}
