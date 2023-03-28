package main

func uniquePaths(m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	dp[0][0] = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j != 0 {
				dp[i][j] = 1
				continue
			}
			if j == 0 {
				dp[i][j] = 1
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func main() {
	uniquePaths(3, 3)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev * ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

