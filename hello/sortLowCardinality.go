package hello

import (
	"fmt"
)

func SortArrayWithLowCardinality(in []int) []int {
	last1, first3 := 0, len(in)-1
	i := 0
	for i <= first3 {
		v := in[i]
		switch v {
		case 1:
			in[i], in[last1] = in[last1], in[i]
			last1++
			i++
		case 2:
			i++
		case 3:
			in[i], in[first3] = in[first3], in[i]
			first3--
		default:
			panic("Invalid number!")
		}
	}
	return in
}

func TestSortArrayWithLowCardinality() {
	fmt.Println(SortArrayWithLowCardinality([]int{3, 2, 1}))
	fmt.Println(SortArrayWithLowCardinality([]int{3, 2, 1, 2, 3}))
	fmt.Println(SortArrayWithLowCardinality([]int{3, 2, 1, 2}))
	fmt.Println(SortArrayWithLowCardinality([]int{1, 2, 3}))
	fmt.Println(SortArrayWithLowCardinality([]int{3, 3, 3, 2, 1}))
	fmt.Println(SortArrayWithLowCardinality([]int{3, 2, 1, 2}))
}
