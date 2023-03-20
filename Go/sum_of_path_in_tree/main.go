package main

import "fmt"

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

func checkPath(root *TreeNode, targetSum int, current *[]int, total *[][]int) {
	if root == nil {
		return
	}
	next := targetSum - root.Val
	if root.Left == nil && root.Right == nil {
		if next == 0 {
			*current = append(*current, root.Val)
			tmp := make([]int, len(*current))
			copy(tmp, *current)
			*total = append(*total, tmp)
			*current = (*current)[:len(*current)-1]
		}
		return
	}
	*current = append(*current, root.Val)
	checkPath(root.Left, next, current, total)
	checkPath(root.Right, next, current, total)
	*current = (*current)[:len(*current)-1]
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var current []int
	var total [][]int
	checkPath(root, targetSum, &current, &total)
	return total
}

func main() {
	oneL := &TreeNode{Val: 1}
	twoL := &TreeNode{Val: 2}
	threeL := &TreeNode{Val: 3}
	oneL.Left = twoL
	oneL.Right = threeL

	oneR := &TreeNode{Val: 2}
	twoR := &TreeNode{Val: 1}
	threeR := &TreeNode{Val: 6}
	oneR.Left = threeR
	oneR.Right = twoR

	root := &TreeNode{Val: 0}
	root.Left = oneL
	root.Right = oneR
	fmt.Println(pathSum(root, 3))
}
