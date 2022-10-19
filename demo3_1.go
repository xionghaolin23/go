package main

import "fmt"

func main() {

	//流程控制 if else if  else

	var x int = 10

	if x > 20 {
		fmt.Println(1)
	} else if x == 10 {
		fmt.Println(2)
	} else {
		fmt.Println(3)
	}

	//switch

	var score int = 90

	switch score {

	case 90:
		fmt.Println("优秀")
	case 80:
		fmt.Println("良好")
	case 70, 60, 50:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")

	}

	//3、fallthrough 穿透  breack 终止穿透

	switch {

	case true:
		fmt.Println(1)
		fallthrough //穿透 会自动执行下面的语句
	case false:
		//break
		fmt.Println(2)
	default:
		fmt.Println(3)

	}

}
