package hello

import (
	"fmt"
	"sort"
)

type state int
type direction int

const Up state = 0
const Left state = 1
const Right state = 2
const LeftDir direction = 0
const RightDir direction = 1

func setStateEntries(stopIndex, startIndex int, result []state, st state) {
	for j := stopIndex; j <= startIndex; j++ {
		result[j] = st
	}
}

func PushDominos(n int, forces [][]int) []state {
	result := make([]state, n)
	sort.Slice(forces, func(a, b int) bool {
		return forces[a][0] < forces[b][0]
	})

	stopIndex := 0
	for i := 0; i < len(forces); i++ {
		force := forces[i]
		startIndex, dir := force[0], force[1]
		switch direction(dir) {
		case LeftDir:
			setStateEntries(stopIndex, startIndex, result, Left)
			stopIndex = startIndex + 1
		case RightDir:
			if len(forces) == i+1 {
				setStateEntries(startIndex, n-1, result, Right)
				break
			}
			nextForce := forces[i+1]
			nextStartIndex, nextDirection := nextForce[0], nextForce[1]
			if direction(nextDirection) == RightDir {
				setStateEntries(startIndex, nextStartIndex-1, result, Right)
				break
			}

			window := nextStartIndex - startIndex
			stopIndex = startIndex + window/2 + 1
			setStateEntries(startIndex, stopIndex, result, Right)
			if window%2 == 0 {
				result[stopIndex-1] = Up
			}
		}
	}
	return result
}

func TestPushDominos() {
	fmt.Println(PushDominos(20, [][]int{{1, int(LeftDir)}, {3, int(RightDir)}, {5, int(RightDir)}, {9, int(LeftDir)}, {12, int(RightDir)}, {15, int(LeftDir)}}))
}
