package main
import (
	"fmt"
)

func getMaxAndIndex(nums *[5]int) (int, int) {
	var max, index = nums[0], 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			index = i
			max = nums[i]
		}
	}
	return max, index
}
func main() {
	var array = [...]int {10, 3, -5, 28, 21}
	max, index := getMaxAndIndex(&array)
	fmt.Printf("最大值是%d, 下标是%d", max, index)
}