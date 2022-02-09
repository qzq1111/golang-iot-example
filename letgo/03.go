package letgo

import (
	"math"
)

/*
给定一个字符串 s ，请你找出其中不含有重复字符的最长子串的长度。

示例1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。
示例 4:

输入: s = ""
输出: 0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func LengthOfLongestSubstring(s string) int {
	temp := make([]int, 128)
	ans := 0
	start := 0

	// 窗口滑动，每次判断当前值是否存在于列表，如果存在判断其位置，然后记录新的位置，然后计算最新长度。 当前位置减去开始位置（重复字符位置）
	for end := 0; end < len(s); end++ {
		start = int(math.Max(float64(start), float64(temp[s[end]])))
		temp[s[end]] = end + 1
		ans = int(math.Max(float64(ans), float64(end-start+1)))
		//fmt.Println(s[end], start, end, ans)
	}
	return ans
}
