package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p != nil && q == nil) ||
		(p == nil && q != nil) {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	oneL := &TreeNode{Val: 1}
	twoL := &TreeNode{Val: 2}
	threeL := &TreeNode{Val: 3}
	oneL.Left = twoL
	oneL.Right = threeL

	oneR := &TreeNode{Val: 1}
	twoR := &TreeNode{Val: 2}
	threeR := &TreeNode{Val: 4}
	oneR.Left = twoR
	oneR.Right = threeR

	fmt.Println(isSameTree(oneL, oneR))
}
