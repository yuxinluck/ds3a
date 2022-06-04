package main

/*
1071. 字符串的最大公因子
对于字符串 s 和 t，只有在 s = t + ... + t（t 自身连接 1 次或多次）时，我们才认定 “t 能除尽 s”。

给定两个字符串 str1 和 str2 。返回 最长字符串 x，要求满足 x 能除尽 str1 且 X 能除尽 str2 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/greatest-common-divisor-of-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func gcdOfStrings(str1 string, str2 string) string {

	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}

	remainder := modOfStr(str1, str2)

	if remainder == "" {
		return str2
	} else if remainder == str1 {
		return ""
	}

	return gcdOfStrings(str2, remainder)

}

//len(str2)<=len(str1)
func modOfStr(str1 string, str2 string) string {
	len2 := len(str2)
	remainder := str1

	for {
		if len(remainder) < len2 || remainder[:len2] != str2 {
			break
		}
		remainder = remainder[len2:]
	}

	return remainder
}

/*
解题思路
采用欧几里得求解最大公因子的方法，来求解字符串的最大公因子。辗转相除法-维基百科
先构建字符串的取余方法，可求得一个字符串对另外一个字符串取余后的余数remainder。
递归地实现辗转相除法，递归基为
余数remainder == ""， 那么返回gcd的最大公因子为str2；
remainder == str1，说明str1无法对str2取余，gcd返回""空字符串。
递归入口：传入先前的str2和str1、str2取余后所得remainder作为参数，进行递归求解。


作者：wenwenwu
链接：https://leetcode.cn/problems/greatest-common-divisor-of-strings/solution/zhan-zhuan-xiang-chu-fa-qiu-jie-zi-fu-chuan-de-zui/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
