package main

import "fmt"

// bubleSort
func BubbleSort(arr []int) []int {
	for i := range arr {
		for j := 0; j < len(arr)-1; j++ {
			if arr[i] < arr[j] {
				arr[i] = arr[i] + arr[j]
				arr[j] = arr[i] - arr[j]
				arr[i] = arr[i] - arr[j]
			}
		}
	}
	return arr
}

// mergeSort
func merge(left []int, right []int) []int {
	var mergeArr []int
	for {
		if len(left) == 0 || len(right) == 0 {
			break
		}
		if left[0] < right[0] {
			mergeArr = append(mergeArr, left[0])
			left = left[1:]
		} else {
			mergeArr = append(mergeArr, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		for i := range left {
			mergeArr = append(mergeArr, left[i])
		}
	}
	if len(right) > 0 {
		for i := range right {
			mergeArr = append(mergeArr, right[i])
		}
	}
	return mergeArr
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {
		return merge(mergeSort(arr[:len(arr)/2]), mergeSort(arr[len(arr)/2:]))
	}
}

/* // Quicksort code loi chua kip sua
func quickSort(arr []int) []int {
	if len(arr) > 2 {
		pivot := arr[len(arr)-1]
		l := 0
		r := len(arr) - 2
		for {
			if l >= r {
				break
			}
			for {
				if arr[l] > pivot || l > r {
					break
				}
				l++
			}
			for {
				if arr[r] <= pivot || l > r {
					break
				}
				r++
			}
			if l >= r {
				break
			}
			arr[l] += arr[r]
			arr[r] = arr[l] - arr[r]
			arr[l] = arr[l] - arr[r]
			l++
			r++
		}
		arr[l] += arr[len(arr)-1]
		arr[len(arr)-1] = arr[l] - arr[len(arr)-1]
		arr[l] = arr[l] - arr[len(arr)-1]
	}
	return arr
} */

func quickSort(arr []int, l int, h int) []int {
	fmt.Println(arr)
	if l < h {
		k := arr[h]
		left := l
		right := h - 1
		for {
			for {
				if arr[left] > k || left >= right {
					break
				}
				left++
			}
			for {
				if arr[right] < k || left >= right {
					break
				}
				right--
			}

			if left >= right {
				break
			}

			arr[left] = arr[left] + arr[right]
			arr[right] = arr[left] - arr[right]
			arr[left] = arr[left] - arr[right]
			left++
			right--
		}
		if arr[left] > arr[h] {
			arr[left] = arr[left] + arr[h]
			arr[h] = arr[left] - arr[h]
			arr[left] = arr[left] - arr[h]
		}
		quickSort(arr, l, left)
		quickSort(arr, left+1, h)
	}
	return arr
}

//
func main() {
	arr := []int{12, 3, 1, 4, 6, 7, 12, 56, 2, 7, 5}
	fmt.Println(quickSort(arr, 0, len(arr)-1))
}
