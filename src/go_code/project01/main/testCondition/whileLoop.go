package main

import "fmt"
func main() {
	var a int
	fmt.Println("请输入一个数字")
	fmt.Scan(&a)
	// if(a == 1){
	// 	fmt.Println("输入正确,值为",a)
	// }else{
	// 	fmt.Println("输入错误,值为",a)
	// }
	// switch {
	// 	case 1 <= a : fmt.Println("您的分数是90分")
	// 		fmt.Println("您的等级是A")
	// 		fallthrough
	// 	case 2 <= a : fmt.Println("您的分数是80分")
	// 		fmt.Println("您的等级是B")
	// 		fallthrough
	// 	default : fmt.Println("您的分数是60分")
	// 		fmt.Println("您刚好及格，等级为C")
	// }
	switch a {
		case 1 : fmt.Println("您的分数是90分")
			fmt.Println("您的等级是A")
			fallthrough
		case 2 : fmt.Println("您的分数是80分")
			fmt.Println("您的等级是B")
			fallthrough
		default : fmt.Println("您的分数是60分")
			fmt.Println("您刚好及格，等级为C")
	}
}