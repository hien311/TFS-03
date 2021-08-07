package main

import "fmt"

func fi(n int) int64 {
	x := int64(1)
	y := int64(1)
	t := x
	for i := 1; i < n; i++ {
		y = y + x
		x = y - x
		t = x
	}
	return t

}

func main() {
	fmt.Println(fi(6))
}
