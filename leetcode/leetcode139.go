package leetcode

import "strings"

/* 背包问题：
	01 背包问题：
     如果是 01 背包，即数组中的元素不可重复使用，外循环遍历 arrs，内循环遍历 target，且内循环倒序:

	完全背包问题：
	（1）如果是完全背包，即数组中的元素可重复使用并且不考虑元素之间顺序，arrs 放在外循环（保证 arrs 按顺序），target在内循环。且内循环正序。
	（2）如果组合问题需考虑元素之间的顺序，需将 target 放在外循环，将 arrs 放在内循环，且内循环正序。
*/
// 从头开始查找的思路
func wordBreak(s string, wordDict []string) bool {
	if s == "" {
		return true
	}
	for _, str := range wordDict {
		if strings.HasPrefix(s, str) {
			ret := wordBreak(s[len(str):], wordDict)
			if ret {
				return true
			}
		}
	}

	return false
}

// 看答案：
func wordBreak2(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, v := range wordDict {
		wordDictSet[v] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break //符合要求
			}
		}
	}
	return dp[len(s)]
}
