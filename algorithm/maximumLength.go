package algorithm

import (
	"slices"
)

/**
 * 描述：给你一个仅由小写英文字母组成的字符串 s 。
 *      如果一个字符串仅由单一字符组成，那么它被称为 特殊 字符串。例如，字符串 "abc" 不是特殊字符串，
 *      而字符串 "ddd"、"zz" 和 "f" 是特殊字符串。
 *      返回在 s 中出现 至少三次 的 最长特殊子字符串 的长度，如果不存在出现至少三次的特殊子字符串，则返回 -1 。
 *      子字符串 是字符串中的一个连续 非空 字符序列。
 *
 * 输入：s = "aaaa"
 * 输出：2
 * 解释：出现三次的最长特殊子字符串是 "aa" ：子字符串 "aaaa"、"aaaa" 和 "aaaa"。可以证明最大长度是 2 。
 *
 * 输入：s = "abcdef"
 * 输出：-1
 * 解释：不存在出现至少三次的特殊子字符串。因此返回 -1 。
 *
 * 输入：s = "abcaba"
 * 输出：1
 * 解释：出现三次的最长特殊子字符串是 "a" ：子字符串 "abcaba"、"abcaba" 和 "abcaba"。可以证明最大长度是 1 。
 */
func MaximumLength(s string) int {
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	// 记录每个特殊子串的长度集合
	groups := [26][]int{}
	count := 0
	for i := 0; i < len(s); i++ {
		count++
		// 最后一个字符或者当前字符与下一个字符不相同，记录特殊字符长度
		if i+1 == len(s) || s[i] != s[i+1] {
			groups[s[i]-'a'] = append(groups[s[i]-'a'], count)
			count = 0
		}
	}

	res := 0
	// 遍历找出最长特殊子串
	for _, group := range groups {
		l := len(group)
		if l == 0 {
			continue
		}
		// 排序，从大到小
		slices.SortFunc(group, func(a, b int) int {
			return b - a
		})
		// 如果最长子串大于2，则特殊子串最小也是group[0]-2
		if l > 0 && group[0] > 2 {
			res = max(res, group[0]-2)
		}
		// 如果至少有两个特殊子串，最大长度的最小值应该是group[0]-1与group[1]中较小的那一个
		if l > 1 {
			temp := min(group[0]-1, group[1])
			res = max(res, temp)
		}
		// 如果有至少3个特殊子串，最大长度的最小值至少是group[2]
		if l > 2 {
			res = max(res, group[2])
		}
	}

	if res == 0 {
		res = -1
	}

	return res
}
