package main

import "fmt"

func main() {

	fmt.Println(1)
	//defer 延迟函数 主要用于一些资源关闭的操作
	defer num(2)
	fmt.Println(3)

	//2、函数的本质本身是一个变量 fun()类型的

	fmt.Printf("%T", num) //打印类型 func(int)

	//既然函数本身是一个变量 那函数也可以给函数赋制

}

func num(i int) {

	fmt.Println(i)

}
