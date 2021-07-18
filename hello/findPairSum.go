package hello

import "fmt"

func FindPairSum(in []int, target int) []int {
	comps := map[int]int{}
	for i, v := range in {
		if otherI, ok := comps[v]; ok {
			return []int{otherI, i}
		}

		comps[target-v] = i
	}
	return nil
}

func TestFindPairSum() {
	fmt.Println(FindPairSum([]int{1, 2, 3, 5, 7}, 8))
	fmt.Println(FindPairSum([]int{1, 2, 3, 5, 7}, 9))
	fmt.Println(FindPairSum([]int{1, 2, 3, 5, 7}, 10))
	fmt.Println(FindPairSum([]int{1, 2, 3, 5, 7}, 11))
	fmt.Println(FindPairSum([]int{1, 2, 3, 5, 7}, 1))
}
