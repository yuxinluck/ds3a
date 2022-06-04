package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
http://www.amoscloud.com/?p=2389

题目描述：
TLV编码是按TagLengthValue格式进行编码的，一段码流中的信元用tag标识，tag在码流中唯一不重复，length表示信元value的长度，value表示信元的值，码流以某信元的tag开头，tag固定占一个字节，length固定占两个字节，字节序为小端序，现给定tlv格式编码的码流以及需要解码的信元tag，请输出该信元的value。

输入码流的16进制字符中，不包括小写字母；
且要求输出的16进制字符串中也不要包含小写字母；
码流字符串的最大长度不超过50000个字节。

输入描述
第一行为第一个字符串 ，表示待解码信元的tag；
输入第二行为一个字符串， 表示待解码的16进制码流；
字节之间用空格分割。

输出描述
输出一个字符串，表示待解码信元以16进制表示的value。

*/

func main() {
	/*
			31
		    32 01 00 AE 90 02 00 01 02 30 03 00 AB 32 31 31 02 00 32 33 33 01 00 CC
	*/

	tag := "31"
	src := "32 01 00 AE 90 02 00 01 02 30 03 00 AB 32 31 31 02 00 32 33 33 01 00 CC"
	str_l := strings.Split(src, " ")

	solution(tag, str_l)

}

func solution(tag string, str_l []string) {

	pos := 0

	for pos < len(str_l) {
		if str_l[pos] == tag {
			lenth := ToInt(str_l[pos+1]) + ToInt(str_l[pos+2])*16
			for i := pos + 2 + 1; i <= pos+2+lenth; i++ {
				fmt.Println(str_l[i])
			}
			return
		} else {
			lenth := ToInt(str_l[pos+1]) + ToInt(str_l[pos+2])*16
			pos = pos + 2 + lenth + 1
		}
	}
}

//将16进制字符串转换为int
func ToInt(s string) int {
	n, _ := strconv.ParseInt("0x"+s, 0, 0)
	return int(n)
}
