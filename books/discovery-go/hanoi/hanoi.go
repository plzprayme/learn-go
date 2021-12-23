package hanoi

import (
	"fmt"
)

func Hanoi(n int) {
	fmt.Println("Number of disks:", n)
	move(n, 1, 2, 3)
}

func move(n, from, to, via int) {
	if n <= 0 {
		return
	}
	move(n-1, from, via, to)
	fmt.Println(from, "->", to)
	move(n-1, via, to, from)
}
