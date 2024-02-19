package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(sottostringaMassima(os.Args[1],os.Args[2]))
}

func sottostringaMassima(s1, s2 string) string {
	len1, len2 := len(s1), len(s2)
	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	lcs := ""
	i, j := len1, len2
	for i > 0 && j > 0 {
		if s1[i-1] == s2[j-1] {
			lcs = string(s1[i-1]) + lcs
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}
	return lcs
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
