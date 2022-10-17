package main

import "fmt"

func main() {

	//变量的定义

	//第一种方式
	var name string = "张三"
	fmt.Printf(name)

	//第二种方式 不显示定义类型
	var name1 = "李四"
	fmt.Printf(name1)

	//第三种方式 编译器自动推测 语法糖
	sex := "男"
	fmt.Println(sex)

	//第四种方式 批量定义
	var (
		age     int    //默认值为0
		address string //默认值为null
	)
	fmt.Println(age, address)

}
