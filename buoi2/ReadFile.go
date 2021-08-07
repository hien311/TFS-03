package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func includeCheck(value string, file []uint8) bool {
	for i := 0; i < len(file)-len(value)-1; i++ {
		count := 0
		for j := range value {
			if value[j] == string(file)[i+j] {
				count++
			}
			if count == len(value) {
				return true
			}
		}
	}
	return false
}

func min(arr []int) int {
	min := arr[0]
	for i := range arr {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func max(arr []int) int {
	max := arr[0]
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func main() {
	// doc toan bo file
	data, _ := ioutil.ReadFile("./data/data.txt")
	// doc 1 phan cua file
	data2, _ := os.Open("./data/data.txt")
	read1 := make([]byte, 5) //doc 5 byte cua chuoi truyen vao read 1
	data2.Read(read1)

	arrString := strings.Split(string(data), ",") // dua string doc duoc thanh mang?
	arrInt := make([]int, len(arrString))         // dua mang string => mang int
	for i := range arrString {
		arrInt[i], _ = strconv.Atoi(arrString[i])
	}
	// kiem tra xem trong file co gia tri cho truoc hay ko
	fmt.Println("8 co trong file: ", string(data), "=========================>", includeCheck("8", data))
	//in ra man hinh
	sort.Ints(arrInt)
	//
	fmt.Println("chuoi da sort: ", arrInt)
	fmt.Println("min: ", min(arrInt), "max: ", max(arrInt))

}
