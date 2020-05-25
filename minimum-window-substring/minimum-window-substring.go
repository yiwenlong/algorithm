package minimum_window_substring

// 76. 最小覆盖子串
// 给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字符的最小子串。
//
// 示例：
//
// 输入: S = "ADOBECODEBANC", T = "ABC"
// 输出: "BANC"
// 说明：
//
// 如果 S 中不存这样的子串，则返回空字符串 ""。
// 如果 S 中存在这样的子串，我们保证它是唯一的答案。

// https://leetcode-cn.com/problems/minimum-window-substring/
func minWindow(s string, t string) string {
	tMap, countMap := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tMap[t[i]] += 1
	}
	left, right, mleft, mL := 0, 0, 0, len(s)+1
	realCount := 0
	for ; right < len(s); right++ {
		ch := s[right]
		if c, ok := tMap[ch]; ok {
			countMap[ch] += 1
			if countMap[ch] <= c {
				realCount++
			}
		}
		if realCount < len(t) {
			continue
		}
		for realCount == len(t) && left <= right {
			l := right - left + 1
			if l < mL {
				mleft, mL = left, l
			}
			ch = s[left]
			if _, ok := tMap[ch]; ok {
				countMap[ch] -= 1
				if countMap[ch] < tMap[ch] {
					realCount -= 1
				}
			}
			left++
		}
	}
	if mL > len(s) {
		return ""
	}
	return s[mleft : mleft+mL]
}

func minWindow2(s string, t string) string {
	sLen, tLen := len(s), len(t)
	if sLen == 0 || tLen == 0 || sLen < tLen {
		return ""
	}
	var i, j, start, found int
	minLen, tCount, sCount := 0x7fffffff, [256]int{}, [256]int{}
	for _, c := range t {
		tCount[c]++
	}
	for j < sLen {
		if found < tLen {
			prev := int(s[j])
			if tCount[prev] > 0 {
				sCount[prev]++
				if sCount[prev] <= tCount[prev] {
					found++
				}
			}
			j++
		}
		for found == tLen {
			if j-i < minLen {
				start, minLen = i, j-i
			}
			next := int(s[i])
			if tCount[next] > 0 {
				sCount[next]--
				if sCount[next] < tCount[next] {
					found--
				}
			}
			i++
		}
	}
	if minLen == 0x7fffffff {
		return ""
	}
	return s[start : start+minLen]
}
