package main

import "fmt"

func main() {

	var x int
	var y float64

	//键盘输入 接受的是变量的地址 取地址用&符号

	fmt.Scan(&x, &y)

	fmt.Println(x, y)

}
