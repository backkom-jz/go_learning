package series

import "fmt"

func GetFibonacci(n int) ([]int, error) {
	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func Squire(n int) int {
	return n * n
}

func init() {
	fmt.Println("series init")
}

func init() {
	fmt.Println("series init2")
}
