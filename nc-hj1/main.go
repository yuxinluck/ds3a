package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
        
		l := strings.Split(in.Text()," ")
		fmt.Println(len(l[len(l)-1]))
	}
}

