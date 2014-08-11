package main

import "fmt"

func power(x, n uint64) uint64 {
	var total uint64 = 1

	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	} else {
		for n > 0 {
			if n & 1 == 1 {
				total *= x
			}
			x *= x
			n = n >> 1
		}
	}

	return total
}


func main() {
	var x, n, result uint64

	fmt.Print("please input x and n:")

	_, err := fmt.Scanf("%d %d", &x, &n)
	if err != nil {
		fmt.Println("error: ", err)
	}

	result = power(x, n)

	if result == 0 {
		fmt.Println("result is to large!")
	} else {
		fmt.Printf("The result is : %d\n", power(x, n))
	}

}

