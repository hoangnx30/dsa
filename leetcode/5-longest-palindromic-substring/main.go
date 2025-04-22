package main

import "fmt"

// Dynamic programming
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	maxLen := 0
	maxStr := s[0:1]

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	for i := range len(s) {
		dp[i][i] = true
		for j := 0; j < i; j++ {
			if s[i] == s[j] && (i-j < 2 || dp[j+1][i-1]) {
				dp[j][i] = true
				fmt.Printf("i: %d, j: %d\n", i, j)

				if i-j+1 > maxLen {
					maxLen = i - j + 1
					maxStr = s[j : i+1]
				}
			}
		}

	}

	fmt.Println(dp)

	return maxStr
}

func expandFromCenter(s string, left int, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return s[left+1 : right]
}

func longestPalindromeExpandFromCenterSolution(s string) string {
	maxLength := 0
	maxStr := s[0:1]

	for i := range s {
		odd := expandFromCenter(s, i, i)
		even := expandFromCenter(s, i, i+1)
		fmt.Println(odd, even)

		if len(odd) > maxLength {
			maxLength = len(odd)
			maxStr = odd
		}

		if len(even) > maxLength {
			maxLength = len(even)
			maxStr = even
		}
	}

	return maxStr
}

func main() {
	str := "babad"

	fmt.Println(longestPalindromeExpandFromCenterSolution(str))
}
