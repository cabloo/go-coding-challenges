package hello

import "fmt"

func ReverseSlice(in []int) {
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
}

func TestReverseSlice() {
	in := []int{1, 2}
	ReverseSlice(in)
	fmt.Println(in)
	in = []int{1, 2, 3, 4, 5}
	ReverseSlice(in)
	fmt.Println(in)
	in = []int{1}
	ReverseSlice(in)
	fmt.Println(in)
	in = []int{}
	ReverseSlice(in)
	fmt.Println(in)
}
