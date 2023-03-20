package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetricRecursive(right *TreeNode, left *TreeNode) bool {
	if right == nil && left == nil {
		return true
	}
	if (right != nil && left == nil) || (right == nil && left != nil) {
		return false
	}
	return right.Val == left.Val && isSymmetricRecursive(left.Left, right.Right) && isSymmetricRecursive(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	return isSymmetricRecursive(root.Left, root.Right)
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
	oneR.Left = threeR
	oneR.Right = twoR

	root := &TreeNode{Val: 0}
	root.Left = oneL
	root.Right = oneR
	fmt.Println(isSymmetric(root))
}
