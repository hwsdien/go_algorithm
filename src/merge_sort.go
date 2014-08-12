package main

import (
	"fmt"
	"time"
	"math/rand"
)

var length int = 10

func sort(arr []int) {
	arr_len := len(arr)

	if arr_len > 1 {
		index := arr_len / 2

		arr1 := arr[:index]
		arr2 := arr[index:arr_len]

		sort(arr1)
		sort(arr2)

		data := merge(arr1, arr2)
		copy(arr, data)
	}
}

func merge(arr1, arr2 []int) []int {
	len1, len2 := len(arr1), len(arr2)
	i, j, k := 0, 0, 0
	data := make([]int, len1 + len2)

	for i < len1 && j < len2 {
		if arr1[i] <= arr2[j] {
			data[k] = arr1[i]
			i++
		} else {
			data[k] = arr2[j]
			j++
		}
		k++
	}

	if i < len1 {
		for _, v := range arr1[i:] {
			data[k] = v
			k++
		}
	}

	if j < len2 {
		for _, v := range arr2[j:] {
			data[k] = v
			k++
		}
	}

	return data
}


func main() {
	arr := make([]int, length)

	rand.Seed(time.Now().UTC().UnixNano())
	for i, _ := range arr {
		arr[i] = rand.Intn(100)
	}

	fmt.Println("data: ", arr)
	sort(arr)
	fmt.Println("sorted data: ", arr)
}
