package hello

import "fmt"

func ArrayProducts(in []int) []int {
	product := 1
	right := make([]int, len(in))
	out := make([]int, len(in))
	for i := len(right) - 1; i >= 0; i-- {
		right[i] = product
		product *= in[i]
	}

	product = 1
	for i, v := range in {
		out[i] = product * right[i]
		product *= v
	}
	return out
}

func TestArrayProducts() {
	fmt.Println(ArrayProducts([]int{1, 2, 3, 4}))
}
