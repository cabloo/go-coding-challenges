package hello

import "fmt"

type Permutation struct {
	orig string
	perm []int
}

func (perm Permutation) SendAllToChannel(out chan string) {
	defer close(out)
	for next := perm.Next(); next != nil; next = perm.Next() {
		out <- string(next)
	}
}

func (perm Permutation) incPerm() {
	for i := len(perm.perm) - 1; i >= 0; i-- {
		if i == 0 || perm.perm[i] < len(perm.perm)-i-1 {
			perm.perm[i]++
			return
		}
		perm.perm[i] = 0
	}
}

func (perm Permutation) Next() []byte {
	defer perm.incPerm()

	if len(perm.orig) == 0 || perm.perm[0] >= len(perm.perm) {
		return nil
	}

	result := []byte(perm.orig)

	for i, v := range perm.perm {
		result[i], result[i+v] = result[i+v], result[i]
	}

	return result
}

func PrintPerms(in string) {
	out := make(chan string, 16)
	perm := Permutation{in, make([]int, len(in))}
	go perm.SendAllToChannel(out)

	for str := range out {
		fmt.Println(str)
	}
}
