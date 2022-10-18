package main

import "fmt"

var userName string = "张三"

func main() {

	//1、变量交换

	var a int = 1

	var b int = 2

	//两个变量的值交换
	a, b = b, a

	fmt.Println(a, b)

	//2、匿名变量 匿名变量的特点是一个下画线“_”，“_”本身就是一个特殊的标识符，被称为空白标识符
	c, _ := test()
	fmt.Println(c)

	//3、变量的作用域 局部变量和全局变量的名字可以相同 优先打印最近的变量
	var userName string = "李四"
	fmt.Printf(userName)

	//4、常量的定义 关键字const
	const name string = "阿狗"
	//name = "这里不能改变，会提示语法报错"
	fmt.Printf(userName)

	//自动推断
	const e, f, g = 11, "字符串", false

	fmt.Println(e, f, g)

	//基本数据类型 布尔 数值（整数、浮点） 字符串
	var flag bool = true
	fmt.Println(flag)

	var number int64 = 1024
	fmt.Println(number)

	var number2 float64 = 3.14
	fmt.Println(number2)

	var sex = "男"
	fmt.Println(sex)

	//类型转换

	var num1 int = 2
	var num2 float64 = 3.15

	var num3 int32

	num3 = int32(num1)

	fmt.Println(num1, num2, num3)

}

func test() (int, int) {
	return 1, 2
}
