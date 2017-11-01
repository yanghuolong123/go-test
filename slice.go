package main

import (
	"fmt"
)

func main() {
	var s = []int{1, 2, 3, 4, 5}
	fmt.Println(s[2])
	fmt.Println(s[2:])
	fmt.Println(s[:2])
	fmt.Println("=======================")

	var r = []rune{}
	r = append(r, 48, 49, 50, 65)
	fmt.Println(r)
	for _, s := range r {
		fmt.Println(string(s))
	}

	str := "aaaaaabbbb"
	fmt.Println([]byte(str))
}
