package main

import (
	"fmt"
	"time"
	"math/rand"
)


func sort(arr [10]int) [10]int {
	var temp int
	length := len(arr)
	i := 1

	for i < length {
		j := i - 1
		temp = arr[i]
		for j > -1 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
		i++
	}

	return arr
}


func main() {
	var arr [10]int
	var temp, i int

	i = 0
	temp = len(arr)

	for i < temp {
		rand.Seed(time.Now().UTC().UnixNano())
		arr[i] = rand.Intn(100)
		i ++
	}

	fmt.Println("data: ", arr)
	arr = sort(arr)
	fmt.Println("sorted data: ", arr)


}
