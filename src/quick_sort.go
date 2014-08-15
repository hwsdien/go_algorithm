package main

import (
	"fmt"
	"time"
	"math/rand"
)

func sort(list []int) {
	length := len(list)
	if length < 1 {
		return
	}

	index := length / 2
	pivot, i, j := list[index], 0, length - 1
	for i < j {
		for i < index && list[i] <= pivot {
			i ++
		}
		if i < j {
			list[index] = list[i]
			index = i
		}

		for j > index && list[j] > pivot {
			j --
		}
		if i < j {
			list[index] = list[j]
			index = j
		}
	}

	list[i] = pivot
	sort(list[:i])
	sort(list[i + 1:])
}

func main() {
	list := make([]int, 10)

	rand.Seed(time.Now().UTC().UnixNano())
	for i,_ := range list {
		list[i] = rand.Intn(100)
	}

	fmt.Println("data: ", list)
	sort(list)
	fmt.Println("sorted data: ", list)
}
