package main

import (
	"fmt"
	"time"
	"math/rand"
)

func sort(list []int) {
	length := len(list)
	times := length / 2

	for i := 0; i < times; i ++ {
		for j := 0; j < length - i - 1; j ++ {
			if list[j] > list[j + 1] {
				list[j], list[j + 1] = list[j + 1], list[j]
			}
		}

		for j := length - i - 1; j > i; j -- {
			if list[j] < list[j - 1] {
				list[j], list[j - 1] = list[j - 1], list[j]
			}
		}
	}
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
