package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	src := "ahj///,//asj///"

	re, _ := regexp.Compile("/+")

	s := strings.Trim(strings.ReplaceAll(src, ",", "/"), "/")

	s = re.ReplaceAllString(s, "/")

	fmt.Println("/" + s)
}

/*
   给定一个url前缀和url后缀
   通过,分割 需要将其连接为一个完整的url
   如果前缀结尾和后缀开头都没有/
   需要自动补上/连接符
   如果前缀结尾和后缀开头都为/
   需要自动去重
   约束：
    不用考虑前后缀URL不合法情况

    输入描述
    url前缀(一个长度小于100的字符串)
    url后缀(一个长度小于100的字符串)
    输出描述
    拼接后的url

    一、
    输入
    /acm,/bb
    输出
    /acm/bb

    二、
    输入
    /abc/,/bcd
    输出
    /abc/bcd

    三、
    输入
    /acd,bef
    输出
    /acd/bef

    四、
    输入
    ,
    输出
    /

*/
