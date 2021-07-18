package hello

import (
	"fmt"
)

func CanMakeRansomNote(desired, available string) bool {
	if len(desired) == 0 {
		return true
	}

	if len(desired) > len(available) {
		return false
	}

	totalExpected := len(desired)
	expectedCounts := map[rune]int{}
	for _, char := range desired {
		expectedCounts[char]++
	}

	for _, char := range available {
		expect := expectedCounts[char]
		if expect > 0 {
			expect--
			totalExpected--
		}

		if totalExpected <= 0 {
			return true
		}
	}

	return false
}

func TestMakeRansomNote() {
	fmt.Println(CanMakeRansomNote("abc", "dbca"))
	fmt.Println(CanMakeRansomNote("abce", "dbca"))
	fmt.Println(CanMakeRansomNote("", "dbca"))
	fmt.Println(CanMakeRansomNote("aab", "abc"))
	fmt.Println(CanMakeRansomNote("aab", "abca"))
}
