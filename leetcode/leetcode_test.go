package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	longestPalindrome("adfbg")
}

func TestWordBreak(t *testing.T) {
	fmt.Println("test leetcode-139: 单词拆分")
	s := "cars"
	wordDic := []string{"car", "ca", "rs"}
	ret := wordBreak(s, wordDic)
	//ret = wordBreak2(s, wordDic) // 答案
	fmt.Printf("leetcode-139: 单词拆分。 ret = %t\n", ret)
}

func TestRemoveNthFromEnd(t *testing.T) {
	fmt.Println("test leetcode-19: 一趟遍历链表删除元素")
}

func TestFindTargetSumWays(t *testing.T) {
	fmt.Println("test leetcode-494: 目标和")
}
