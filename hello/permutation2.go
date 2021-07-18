package hello

import "fmt"

type Perm struct {
	orig []int
	perm []int
}

func nextPerm(p Perm) []int {
	for i := len(p.perm) - 1; i >= 0; i-- {
		if i == 0 && p.perm[0] >= len(p.perm)-1 {
			return nil
		}
		if i == 0 || p.perm[i] < len(p.perm)-i-1 {
			p.perm[i]++
			return p.perm
		}
		p.perm[i] = 0
	}
	return p.perm
}

func FindPermutations(in []int) chan []int {
	out := make(chan []int)
	go func() {
		defer close(out)
		if len(in) == 0 {
			return
		}
		p := Perm{in, make([]int, len(in))}
		r := make([]int, len(in))
		copy(r, in)
		out <- in
		for nextPerm(p) != nil {
			r := make([]int, len(in))
			copy(r, in)
			for i, delta := range p.perm {
				r[i], r[i+delta] = r[i+delta], r[i]
			}
			out <- r
		}
	}()
	return out
}

func printPerms(in []int) {
	for str := range FindPermutations(in) {
		fmt.Println(str)
	}
}

func TestPerm() {
	printPerms([]int{1})
	printPerms([]int{1, 2, 3})
	printPerms([]int{})
	printPerms([]int{5})
}
