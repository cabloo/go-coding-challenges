package hello

import "fmt"

func findFirstIndex(arr []int, target int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		min, max := 0, len(arr)-1
		for min < max {
			mid := int((max-min)/2 + min)
			v := arr[mid]
			if v < target {
				min = mid + 1
				continue
			}

			if v > target {
				max = mid - 1
				continue
			}

			max = mid
			if max-min <= 1 {
				break
			}
		}

		if arr[min] == target {
			out <- min
		}
	}()
	return out
}
func findLastIndex(arr []int, target int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		min, max := 0, len(arr)-1
		for min < max {
			mid := int((max-min)/2 + min)
			v := arr[mid]
			if v < target {
				min = mid + 1
				continue
			}

			if v > target {
				max = mid - 1
				continue
			}

			min = mid
			if max-min == 1 {
				if arr[max] == target {
					min = max
				}
				break
			}
		}
		if arr[min] == target {
			out <- min
		}
	}()
	return out
}

func FindIndicesInSortedArray(arr []int, target int) []int {
	first := findFirstIndex(arr, target)
	last := findLastIndex(arr, target)
	f, ok := <-first
	l := <-last
	if !ok {
		return []int{}
	}
	return []int{f, l}
}

func TestFindIndicesInSortedArray() {
	fmt.Println(FindIndicesInSortedArray([]int{1, 2, 7, 9, 9, 9, 10}, 9))
	fmt.Println(FindIndicesInSortedArray([]int{1, 2, 7, 9, 9, 9, 10}, 7))
	fmt.Println(FindIndicesInSortedArray([]int{1, 2, 7, 9, 9, 9, 10}, 8))
	fmt.Println(FindIndicesInSortedArray([]int{1, 2, 7, 9, 9, 9, 10}, 1))
	fmt.Println(FindIndicesInSortedArray([]int{1, 2, 7, 9, 9, 9, 10}, 0))
	fmt.Println(FindIndicesInSortedArray([]int{9, 9, 9, 9, 9, 9, 9}, 9))
	fmt.Println(FindIndicesInSortedArray([]int{9}, 9))
	fmt.Println(FindIndicesInSortedArray([]int{9, 9}, 9))
}
