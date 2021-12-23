package main

import (
	"discovery-go/hanoi"
	"discovery-go/seq"
	"discovery-go/strings"
	"fmt"
)

func main() {
	hanoi.Hanoi(3)
	fmt.Println(seq.Fib(6))
	strings.Unicode("가갛힣")
}
