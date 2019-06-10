package main
import "fmt"

func main() {
	// var name string
	// var age byte
	// var score float32
	// var isPass bool
	// // fmt.Println("请输入姓名")
	// // fmt.Scanln(&name)
	// // fmt.Println("请输入年龄")
	// // fmt.Scanln(&age)
	// // fmt.Println("请输入分数")
	// // fmt.Scanln(&score)
	// // fmt.Println("请输入是否通过")
	// // fmt.Scanln(&isPass)
	// // fmt.Println(name, age, score, isPass)
	// fmt.Scanf("%s %d %f %t", &name, &age, &score, &isPass)
	// fmt.Println(name, age, score, isPass)
	fmt.Println("hello,World")
	label1:
	for index := 0; index < 10; index++ {
		fmt.Println("index = ", index)
		for i :=0; i < 10; i++ {
			if i == 2 {
				continue label1
			}
			fmt.Println("i = ", i)
			fmt.Println(hello(i, index))
		}
	}
	fmt.Println("hello，你好")
}
func hello(a int, b int) int{
	return a + b;
}