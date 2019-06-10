package main

import (
	"fmt"
)
type Box struct{
	Length float64
	Weight float64
	Height float64
}

func (box *Box) GetVolumn() float64{
	return box.Length * box.Weight * box.Height
}

func main() {
	var length, weight, height float64
	fmt.Println("请输入长方体的长宽高，以空格隔开")
	fmt.Scanf("%f %f %f", &length, &weight, &height)
	var box = &Box{
		Length : length,
		Weight : weight,
		Height : height,
	}
	volumn := box.GetVolumn()
	fmt.Println("长方体Box的体积是", volumn)
}