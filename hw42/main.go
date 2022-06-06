package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	src := "20-21,15,18,30,5-10"
	ask := 15

	m := make(map[int]bool)
	//反序列化vlan
	for _, vlan := range strings.Split(src, ",") {
		if strings.Contains(vlan, "-") {
			t1 := strings.Split(vlan, "-")[0]
			t2 := strings.Split(vlan, "-")[1]
			start, _ := strconv.Atoi(t1)
			end, _ := strconv.Atoi(t2)
			for i := start; i <= end; i++ {
				m[i] = true
			}
			continue
		}

		t, _ := strconv.Atoi(vlan)
		m[t] = true
	}

	fmt.Println(m)
	//删除申请的
	delete(m, ask)

	vlans := make([]int, 0, len(m))
	for k := range m {
		vlans = append(vlans, k)
	}

	sort.Slice(vlans, func(i, j int) bool {
		return vlans[i] < vlans[j]
	})

	fmt.Println(vlans)

	//序列话
	ans := ""
	start := vlans[0]
	pre := start

	for i := 1; i < len(vlans); i++ {
		cur := vlans[i]
		if cur == pre+1 {
			pre = cur
		} else {
			if start == pre {
				ans = ans + strconv.Itoa(start) + ","
			} else {
				ans = ans + strconv.Itoa(start) + "-" + strconv.Itoa(pre) + ","
			}
			start = cur
			pre = cur
		}
	}

	if start == pre {
		ans = ans + strconv.Itoa(start)
	} else {
		ans = ans + strconv.Itoa(start) + "-" + strconv.Itoa(pre)
	}

	fmt.Println(ans)

}
