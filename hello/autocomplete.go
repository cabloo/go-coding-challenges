package hello

import (
	"fmt"
	"unicode/utf8"
)

type TrieNode struct {
	isWord   bool
	children map[rune]*TrieNode
}

func NewTrie() *TrieNode {
	return &TrieNode{isWord: false, children: map[rune]*TrieNode{}}
}

func (t *TrieNode) AddWord(word string) {
	if len(word) == 0 {
		return
	}

	firstChar, width := utf8.DecodeRuneInString(word)
	if width == 0 {
		return
	}
	rest, ok := t.children[firstChar]
	if !ok {
		rest = NewTrie()
		t.children[firstChar] = rest
	}
	if len(word) == width {
		rest.isWord = true
		return
	}

	rest.AddWord(word[width:])
}

func (t *TrieNode) FinishChoices(start string) chan string {
	result := make(chan string)
	go func() {
		defer close(result)
		n := t
		currSearch := start
		for len(currSearch) > 0 {
			firstChar, width := utf8.DecodeRuneInString(currSearch)
			if width == 0 {
				return
			}
			var ok bool
			n, ok = n.children[firstChar]
			if !ok {
				return
			}
			if len(currSearch) == width {
				break
			}
			currSearch = currSearch[width:]
		}

		n.SendAllWordsToChannelWithPrefix(result, start)
	}()
	return result
}

func (t *TrieNode) SendAllWordsToChannelWithPrefix(out chan string, prefix string) {
	if t.isWord {
		out <- prefix
	}
	for r, n := range t.children {
		n.SendAllWordsToChannelWithPrefix(out, prefix+string(r))
	}
}

func printChan(c chan string) {
	for s := range c {
		fmt.Println(s)
	}
}

func printFinishChoices(t *TrieNode, in string) {
	fmt.Println("\nChecking '", in, "'")
	printChan(t.FinishChoices(in))
}

func TestFinishChoices() {
	t := NewTrie()
	t.AddWord("Zane")
	t.AddWord("Hooper")
	t.AddWord("Zack")
	t.AddWord("Zach")

	printFinishChoices(t, "")
	printFinishChoices(t, "Hello")
	printFinishChoices(t, "?")
	printFinishChoices(t, "Z")
	printFinishChoices(t, "Za")
	printFinishChoices(t, "Zac")
	printFinishChoices(t, "Zan")
	printFinishChoices(t, "Zane")
	printFinishChoices(t, "H")
	printFinishChoices(t, "Hoop")
	printFinishChoices(t, "Hooper")
}
