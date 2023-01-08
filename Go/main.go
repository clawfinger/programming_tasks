package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	a := []int32{1, 2, 5, 3, 7}
	b := a[1:3]
	fmt.Println(a)
	fmt.Println(b)
	b = append(b, 5)
	fmt.Println(a)
	fmt.Println(b)
}
