package main

import "fmt"

func main() {

	//遍历string

	var name string = "xionghaolin"

	fmt.Println(name)

	//查看字符串的长度

	fmt.Println(len(name))

	//第一种遍历字符串的长度
	for i := 0; i < len(name); i++ {
		fmt.Printf("%c", name[i])
	}

	//遍历字符串 i下表 item元素 需要%c展示
	for i, item := range name {
		fmt.Print(i)
		fmt.Printf("%c", item)
	}

}
