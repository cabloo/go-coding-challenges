package hello

import (
	"fmt"
	"math"
)

func FindPythagTriplets(in []int) [][]int {
	out := [][]int{}
	exists := map[int]struct{}{}
	for _, v := range in {
		exists[v*v] = struct{}{}
	}

	for i, v1 := range in {
		sq1 := v1 * v1
		for j := i + 1; j < len(in); j++ {
			v2 := in[j]
			sq2 := v2 * v2
			target := sq1 + sq2

			if _, ok := exists[target]; ok {
				out = append(out, []int{v1, v2, int(math.Sqrt(float64(target)))})
			}
		}
	}
	return out
}

func TestPythagTriplets() {
	fmt.Println(FindPythagTriplets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}))
}
