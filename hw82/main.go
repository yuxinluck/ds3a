package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
http://www.amoscloud.com/?p=3504

1.输入字符串s输出s中包含所有整数的最小和，
  说明：1字符串s只包含a~z,A~Z,+,-，
  2.合法的整数包括正整数，一个或者多个0-9组成，如：0,2,3,002,102
  3.负整数，负号开头，数字部分由一个或者多个0-9组成，如-2,-012,-23,-00023
  输入描述：包含数字的字符串
  输出描述：所有整数的最小和
  示例：
    输入：
      bb1234aa
  　输出
      10
  　输入：
      bb12-34aa
  　输出：
      -31
  说明：1+2-(34)=-31

*/

func main() {

	s := "bb12-34aa"

	l := strings.Split(s, "")

	sum := 0

	for i := 0; i < len(l); i++ {

		c := l[i]

		if c == "-" {
			i++
			nums := make([]string,0)

			for i < len(l) && isNum(l[i]) {
				nums = append(nums, l[i])
				i++
			}
			
			if len(nums) >0{
                t,_ := strconv.Atoi(strings.Join(nums,""))
                sum -= t
			}
			i--
			continue
		}

		if isNum(c) {
			t,_ := strconv.Atoi(c)
			sum += t
		}
	}
	fmt.Println(sum)

}

func isNum(s string) bool {
	_, err := strconv.ParseInt(s, 0, 0)
	return err == nil
}
