package blind75

func countBits(n int) []int {
	dp := make([]int, n+1)
	offset := 1
	for i := 1; i < len(dp); i++ {
		if offset*2 == i {
			offset = i
		}
		dp[i] = 1 + dp[i-offset]
	}
	return dp
}
