package main

import (
	"fmt"
)

func main() {
	var s = []int{1, 2, 3, 4, 5}
	fmt.Println(s[2])
	fmt.Println(s[2:])
	fmt.Println(s[:2])
}
