package hello

import (
	"fmt"
)

func FindNonDuplicate(in []int) int {
	seen := map[int]struct{}{}
	seenOnce := map[int]struct{}{}
	for _, v := range in {
		_, seenM := seen[v]
		if seenM {
			continue
		}
		_, seenO := seenOnce[v]
		if seenO {
			delete(seenOnce, v)
			seen[v] = struct{}{}
			continue
		}
		seenOnce[v] = struct{}{}
	}

	for k := range seenOnce {
		return k
	}
	return 0
}

func TestFindNonDuplicate() {
	fmt.Println(FindNonDuplicate([]int{1, 2, 3, 2, 3}))
	fmt.Println(FindNonDuplicate([]int{1, 2, 2, 3, 2, 2, 3, 3}))
	fmt.Println(FindNonDuplicate([]int{2, 2, 3, 4, 3}))
}
