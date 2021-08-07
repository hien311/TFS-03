package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func primeNumberCheck(p int64) bool {
	var y float64 = float64(p)
	if p < 2 {
		return false
	} else if p == 2 || p == 3 {
		return true
	} else {
		var i int64
		for i = 2; i <= int64(math.Sqrt(y)); i++ {
			if p%i == 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	data, _ := ioutil.ReadFile("./data/primesdata.txt")
	s := strings.Split(string(data), ",")
	arr := make([]int64, len(s))
	for i := range s {
		arr[i], _ = strconv.ParseInt(s[i], 10, 64)
		fmt.Println(primeNumberCheck(arr[i]))
	}

}
