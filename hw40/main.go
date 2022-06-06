package main

import (
	"fmt"
	"strings"
)

func main() {
	m := make(map[bool][]string)
	m[true] = []string{}
	m[false] = []string{}

	src := "1/N 2/Y 3/N 4/Y"

	stus := strings.Split(src, " ")
	fmt.Println(stus)

	f := true
	for _, s := range stus {

		if strings.HasSuffix(s, "N") {
			f = !f
		}
		m[f] = append(m[f], s)
	}

	for _, v := range m {
		fmt.Println(v)
	}
}
