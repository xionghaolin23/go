package main

import "fmt"

func main() {

	var sum int = 0

	//for 循环
	for i := 1; i <= 5; i++ {
		sum = sum + i
	}
	fmt.Println(sum)

}
