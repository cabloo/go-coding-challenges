package hello

import (
	"fmt"
	"math"
)

type Solution struct {
	// a is always at least as long as b
	a, b []int
	i, j int
}

func (s *Solution) findMedian() float64 {
	if len(s.a) == 0 {
		return median(s.b)
	}
	if len(s.b) == 0 {
		return median(s.a)
	}

	for !s.hasReachedTarget() {
		s.moveTowardsTarget()
	}

	if (len(s.a)+len(s.b))%2 == 1 {
		if s.countLeft() > s.countRight() {
			return float64(s.maxLeft())
		}
		return float64(s.minRight())
	}

	return float64(s.maxLeft()+s.minRight()) / 2
}

func (s *Solution) maxLeft() int {
	if s.i < 0 {
		return s.b[len(s.b)-1]
	}
	return max(s.maxLeftA(), s.maxLeftB())
}

func (s *Solution) minRight() int {
	if s.i < -1 {
		return s.a[0]
	}
	return min(s.minRightA(), s.minRightB())
}

func (s *Solution) moveTowardsTarget() {
	dist := 1 + min(len(s.b)-s.j, s.j)/2
	if s.maxLeftB() > s.minRightA() {
		dist *= -1
	}
	s.i -= dist
	s.j += dist
}

func (s *Solution) countLeft() int {
	return min(max(s.i+1, 0), len(s.a)) + min(max(s.j+1, 0), len(s.b))
}

func (s *Solution) countRight() int {
	return len(s.a) - min(max(s.i+1, 0), len(s.a)) + len(s.b) - min(max(s.j+1, 0), len(s.b))
}

func (s *Solution) maxLeftA() int {
	if s.i >= len(s.a) || s.i < 0 {
		return math.MinInt32
	}
	return s.a[s.i]
}

func (s *Solution) maxLeftB() int {
	if s.j >= len(s.b) || s.j < 0 {
		return math.MinInt32
	}
	return s.b[s.j]
}

func (s *Solution) minRightA() int {
	if s.i+1 >= len(s.a) || s.i < -1 {
		return math.MaxInt32
	}
	return s.a[s.i+1]
}

func (s *Solution) minRightB() int {
	if s.j+1 >= len(s.b) || s.j < -1 {
		return math.MaxInt32
	}
	return s.b[s.j+1]
}

func (s *Solution) hasReachedTarget() bool {
	return s.maxLeft() <= s.minRight()
}

func NewSolution(nums1, nums2 []int) *Solution {
	j := (len(nums2) - 1) / 2
	i := iFromJ(nums1, nums2, j)
	return &Solution{nums1, nums2, i, j}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func iFromJ(nums1, nums2 []int, j int) int {
	return (len(nums1)+len(nums2))/2 - j - 2
}

func median(nums []int) float64 {
	if len(nums)%2 == 1 {
		return float64(nums[len(nums)/2])
	}
	return (float64(nums[len(nums)/2]) + float64(nums[(len(nums)-1)/2])) / 2
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums2) > len(nums1) {
		nums1, nums2 = nums2, nums1
	}
	return NewSolution(nums1, nums2).findMedian()
}

func TestMedianSortedArrays() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{5, 6, 7, 8}, []int{1, 2, 3, 4}))
	// fmt.Println(findMedianSortedArrays([]int{5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4}))
}
