package main

import "fmt"

func longestPalindrome(s string) string {
	res := ""
	length := 0
	for i := 0; i < len(s); i++ {
		l := i
		r := i
		for l >= 0 && r <= len(s)-1 && s[l] == s[r] {
			if r-l+1 > length {
				length = r - l + 1
				res = s[l : r+1]
			}
			l--
			r++
		}
		l = i
		r = i + 1
		for l >= 0 && r <= len(s)-1 && s[l] == s[r] {
			if r-l+1 > length {
				length = r - l + 1
				res = s[l : r+1]
			}
			l--
			r++
		}
	}
	return res
}

func main() {
	fmt.Println(longestPalindrome("abba"))
}
