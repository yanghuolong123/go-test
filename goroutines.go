package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("1111111111")
		go func() {
			fmt.Println("222222222")
		}()
		fmt.Println("333333333")
	}()
	time.Sleep(5 * time.Second)
}
