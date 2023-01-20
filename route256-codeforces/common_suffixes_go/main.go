package main

import (
	"bufio"
	"fmt"
	"os"
)

func substrings(str string, dict map[string][]string) {
	for i := 0; i < len(str); i++ {
		dict[str[i:]] = append(dict[str[i:]], str)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var dictSize int
	fmt.Fscan(in, &dictSize)
	dict := make(map[string][]string)
	for i := 0; i < dictSize; i++ {
		var entry string
		fmt.Fscan(in, &entry)
		substrings(entry, dict)
	}
	var wordsNum int
	fmt.Fscan(in, &wordsNum)
	words := make([]string, 0, wordsNum)
	for i := 0; i < wordsNum; i++ {
		entry := ""
		fmt.Fscan(in, &entry)
		words = append(words, entry)
	}
	for _, word := range words {
		found := false
		for i := 0; i < len(word); i++ {
			suffix := word[i:]
			dictWords, ok := dict[suffix]
			if ok {
				for _, dictWord := range dictWords {
					if dictWord != word {
						fmt.Println(dictWord)
						found = true
						break
					}
				}
			}
			if found {
				break
			}
		}
		if !found {
			for _, words := range dict {
				for _, dictWord := range words {
					if word != dictWord {
						fmt.Println(dictWord)
						goto mark
					}
				}
			}
		mark:
		}
	}
}
