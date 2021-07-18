package hello

import "fmt"

func MergeSort(in []int) []int {
	if len(in) <= 1 {
		result := make([]int, len(in))
		copy(result, in)
		return result
	}

	mid := len(in) / 2
	var left []int
	done := make(chan bool)
	go func() {
		left = MergeSort(in[:mid])
		done <- true
	}()

	right := MergeSort(in[mid:])
	<-done

	return Merge(left, right)
}

func Merge(left, right []int) []int {
	l, r := 0, 0
	result := make([]int, len(left)+len(right))
	for {
		if l < len(left) && (r >= len(right) || left[l] < right[r]) {
			result[l+r] = left[l]
			l++
			continue
		}
		if r < len(right) {
			result[l+r] = right[r]
			r++
			continue
		}
		return result
	}
}

func TestMergeSort() {
	lg := make([]int, 10000000)
	for i := 0; i < len(lg); i++ {
		lg[i] = len(lg) - i
	}
	MergeSort(lg)
	fmt.Println(MergeSort([]int{5, 4, 2, 3, 1}))
	fmt.Println(MergeSort([]int{2, 1}))
	fmt.Println(MergeSort([]int{1}))
	fmt.Println(MergeSort([]int{}))
}
