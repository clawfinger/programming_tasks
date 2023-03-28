package main

import "sort"

//3,2,2 -> 1,1,1
func findContentChildren(g []int, s []int) int {
    sort.Slice(g, func(i int, j int)bool {
		return g[i] > g[j]
	})
	sort.Slice(s, func(i int, j int)bool {
		return s[i] > s[j]
	})
	cookie := 0
	res := 0
	for i:= 0; i < len(g); i++ {
		if cookie == len(s) { //c=0 len(3)
			break
		}
		if g[i] <= s[cookie] { // g=2 c = 1
			res++ //0
			cookie++//0
		}
	}
	return res
}

