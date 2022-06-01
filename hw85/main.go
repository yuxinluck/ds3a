package main

import "fmt"

/*
http://www.amoscloud.com/?p=3559

有一个特殊的五键键盘
    上面有A、Ctrl-C、Ctrl-X、Ctrl-V、Ctrl-A
    A键在屏幕上输出一个字母A
    Ctrl-C将当前所选的字母复制到剪贴板
    Ctrl-X将当前选择的字母复制到剪贴板并清空所选择的字母
    Ctrl-V将当前剪贴板的字母输出到屏幕
    Ctrl-A选择当前屏幕中所有字母
    注意：
      1. 剪贴板初始为空
      2. 新的内容复制到剪贴板会覆盖原有内容
      3. 当屏幕中没有字母时,Ctrl-A无效
      4. 当没有选择字母时Ctrl-C、Ctrl-X无效
      5. 当有字母被选择时A和Ctrl-V这两个输出功能的键,
         会先清空所选的字母再进行输出

    给定一系列键盘输入,
    输出最终屏幕上字母的数量

    输入描述:
       输入为一行
       为简化解析用数字12345分别代替A、Ctrl-C、Ctrl-X、Ctrl-V、Ctrl-A的输入
       数字用空格分割

    输出描述:
        输出一个数字为屏幕上字母的总数量

    示例一:
        输入:
          1 1 1
        输出:
          3

   示例二:
        输入:
          1 1 5 1 5 2 4 4
        输出:
          2

*/

func main() {

	nums := []int{1, 1, 5, 1, 5, 2, 4, 4}
	//选择
	choose := make([]int, 0)
	//剪切版
	clip := make([]int, 0)

	//屏幕
	screen := make([]int, 0)

	//标记是否已经清除
	cleaned := false

	for i := 0; i < len(nums); i++ {
		switch nums[i] {
		case 1: // 会先清空所选的字母再进行输出,A键在屏幕上输出一个字母A
			if len(choose) > 0 && !cleaned {
				if len(screen) == len(choose) {
					screen = []int{}
				} else {
					screen = screen[len(choose):]
				}
				cleaned = true
			}
			screen = append(screen, 1)
			fmt.Println(1, screen)
		case 2: // Ctrl-C将当前所选的字母复制到剪贴板
			if len(choose) > 0 {
				temp := make([]int, len(choose))
				copy(temp, choose)
				clip = temp
			}
			fmt.Println(2, clip)
		case 3: //Ctrl-X将当前选择的字母复制到剪贴板并清空所选择的字母
			if len(choose) > 0 {
				temp := make([]int, len(choose))
				copy(temp, choose)
				clip = temp
				choose = []int{}
			}
		case 4: //会先清空所选的字母再进行输出,Ctrl-V将当前剪贴板的字母输出到屏幕
			if len(choose) > 0 && !cleaned {
				if len(screen) == len(choose) {
					screen = []int{}
				} else {
					screen = screen[len(choose):]
				}
				cleaned = true
			}
			screen = append(screen, clip...)
		case 5: //Ctrl-A选择当前屏幕中所有字母,当屏幕中没有字母时,Ctrl-A无效
			if len(screen) > 0 {
				temp := make([]int, len(screen))
				copy(temp, screen)
				choose = temp
				fmt.Println(5, choose)
				cleaned = false
			}
		}
	}

	fmt.Println(len(screen))
}
