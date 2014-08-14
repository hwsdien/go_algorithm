package main

import (
	"fmt"
	"time"
	"math/rand"
)

func sort(list []int) {
	count_list := make([]int, 100)

	for _,v := range list {
		count_list[v] ++
	}


	length := len(count_list)
	for i := 1; i < length; i++ {
		count_list[i] += count_list[i - 1]
	}

	list_len := len(list)
	sorted_list := make([]int, list_len)
	for j := list_len; j > 0; j -- {
		sorted_list[count_list[list[j - 1]] - 1] = list[j - 1]
		count_list[list[j - 1]] --
	}

	copy(list, sorted_list)

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
