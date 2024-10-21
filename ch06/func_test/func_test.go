package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func Sum(op ...int) int {
	ret := 0
	for _, v := range op {
		ret += v
	}
	return ret
}

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("clear")
	}()
	t.Log("call")
	panic("fatal error")
}

func TestFunc(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)

	myFun := timeSpent(slowFun)
	myFun(10)

	ts := Sum(1, 2, 3, 4, 5)
	fmt.Println(ts)
}
