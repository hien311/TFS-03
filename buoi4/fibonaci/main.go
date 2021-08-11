package main

import (
	"fibonaci/fi"
	"fmt"
)

func main() {
	c := make(chan int64)
	go fi.Fi(50, c)
	for v := range c {
		fmt.Println(v)
	}
}
