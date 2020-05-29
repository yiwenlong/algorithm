package main

import "fmt"

//
//434. 字符串中的单词数
//统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。
//
//请注意，你可以假定字符串里不包括任何不可打印的字符。
//
//示例:
//
//输入: "Hello, my name is John"
//输出: 5
//解释: 这里的单词是指连续的不是空格的字符，所以 "Hello," 算作 1 个单词。
//
func countSegments(s string) int {
	res := 0
	for i := 0; i < len(s); {
		if s[i] != ' ' {
			res += 1
			for i < len(s) && s[i] != ' ' {
				i++
			}
		}
		for i < len(s) && s[i] == ' ' {
			i++
		}
	}
	return res
}

func main() {
	fmt.Printf("res: %d\n", countSegments("Hello, my name is John"))
}
