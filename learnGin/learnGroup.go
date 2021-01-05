package learnGin

import "fmt"

func main() {
	fmt.Println("test go build")
}

type Inter01 interface {
	DDD() string
}

type CC struct {
}

func (c CC) DDD() string {
	return "33"
}

func EE() Inter01 {

	return &CC{}
}
