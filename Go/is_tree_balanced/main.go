package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(height(node.Left), height(node.Right)) + 1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l := height(root.Left)
	r := height(root.Right)
	return math.Abs(float64(l-r)) < 2 && isBalanced(root.Left) && isBalanced(root.Right)
}

func main() {
	oneL := &TreeNode{Val: 1}
	twoL := &TreeNode{Val: 2}
	threeL := &TreeNode{Val: 3}
	oneL.Left = twoL
	oneL.Right = threeL

	oneR := &TreeNode{Val: 1}
	// twoR := &TreeNode{Val: 2}
	// threeR := &TreeNode{Val: 4}
	// oneR.Left = threeR
	// oneR.Right = twoR

	root := &TreeNode{Val: 0}
	root.Left = oneL
	root.Right = oneR
	fmt.Println(isBalanced(root))
}
