package add_strings

//
//给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
//
//注意：
//
//num1 和num2 的长度都小于 5100.
//num1 和num2 都只包含数字 0-9.
//num1 和num2 都不包含任何前导零。
//你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/add-strings
//
// 我的思路：
// 1. 其中最长的最为遍历基准
// 2. 使用一个长度比最长字符串+1的字节数组保存结果
// 3. 计算出数字直接放入到对应的字节数组中, 进位保留，下次操作使用
// 4. 处理结果，如果有进位，则把进位填到结果的0位置，没有，则返回[1:]的数据
func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	i, j := len(num1)-1, len(num2)-1
	var carry, tmp, num uint8
	res := make([]byte, len(num1)+1)
	for {
		if i < 0 {
			break
		}
		if j >= 0 {
			tmp = num1[i] - '0' + num2[j] - '0' + carry
		} else {
			tmp = num1[i] + carry - '0'
		}
		carry = tmp / 10
		num = tmp % 10
		res[i+1] = num + '0'
		i--
		j--
	}
	if carry > 0 {
		res[0] = carry + '0'
		return string(res)
	}
	return string(res[1:])
}
