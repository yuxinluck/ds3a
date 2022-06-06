package main

import (
	"fmt"
	"regexp"

	"strings"
)

var s = `The furthest distance in the world,Is not between life and death,But when I stand in front or you,Yet you don't know that I love you.`

func main() {

	re, _ := regexp.Compile("[^A-Za-z]+")
	words := re.Split(s, -1)
	fmt.Println(words)

	in := "f"

	buf := make([]string, 0)
	for _, v := range words {
		if strings.HasPrefix(v, in) {
			buf = append(buf, v)
		}
	}

	if len(buf) == 0 {
		fmt.Println(in)
	} else {
		for _, v := range buf {
			fmt.Println(v)
		}
	}

}
