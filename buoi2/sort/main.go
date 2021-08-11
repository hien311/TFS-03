package main

import (
	"fmt"
	sort "tfs/buoi2/sort/sortfunc"
)

// bubleSort

//
func main() {
	arr := []int{12, 3, 1, 4, 6, 7, 12, 56, 2, 7, 5}
	fmt.Println(sort.QuickSort(arr, 0, len(arr)-1))
}
