package main

import "fmt"

func main() {

	//函数的使用

	i := add(2, 2)

	fmt.Println(i)

	i2, i3 := swap(1, 2)
	fmt.Println(i2, i3)

	sum := getSum(1, 2, 3)
	fmt.Println(sum)

}

//func 方法名（参数1，参数2) 返回类型
func add(a int, b int) int {
	c := a + b
	return c
}

//多个返回值的函数
func swap(x int, y int) (int, int) {
	return y, x
}

//可变参数
func getSum(nums ...int) int {
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	return sum
}
