package leetcode

import (
	"fmt"
	"strings"
)

func longestPalindrome(s string) string {
	var (
		max string = ""
		l   int    = len(s)
	)
	if l == 0 {
		return ""
	}
	for j := 0; j < l; j++ {
		// 剩余的字符串小于等于找出的最大串就直接返回了
		if len(max) >= (l - j) {
			return max
		}
		last_j := strings.LastIndex(s, string(s[j]))
		if last_j == -1 {
			continue
		} else {
			// 判断是否是回文子串
			for i := 1; i < l; i++ {
				fmt.Println(i, s[i])
			}
		}
	}
	return max
}
