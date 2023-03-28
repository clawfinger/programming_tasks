package main

import "fmt"

func generateRec(s string, n int, open int, closed int) {
	if len(s) == 2*n {
		if open == closed {
			fmt.Println(s)
		}
		return
	}
	generateRec(s+"(", n, open+1, closed)
	if closed < open {
		generateRec(s+")", n, open, closed+1)
	}
}

func generate(n int) {
	var s string
	generateRec(s, 3, 0, 0)
}

func main() {
	generate(3)
}
