package error

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

var LessThanTwoError = errors.New("n should be less than 2")
var LargerThenHundredError = errors.New("n should be larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThenHundredError
	}
	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}
func TestError(t *testing.T) {
	if v, err := GetFibonacci(2); err != nil {
		if errors.Is(err, LessThanTwoError) {
			fmt.Println("It is less.")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func GetFibonacci1(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err == nil {
		if list, err = GetFibonacci(i); err == nil {
			fmt.Println(list)
		} else {
			fmt.Println("Error", err)
		}
	} else {
		fmt.Println("Error", err)
	}
}

func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
		return
	}
	if list, err = GetFibonacci(i); err != nil {

		fmt.Println("Error", err)
		return
	}
	fmt.Println(list)

}
