package dock_test

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}

	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }

	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	mySet[2] = false

	t.Log(mySet)
	// 判断
	if mySet[1] {
		t.Log("1 is Existing !")
	} else {
		t.Log("1 is not Existing !")
	}
	mySet[3] = true
	// 长度
	t.Log(len(mySet))

	// 删除
	delete(mySet, 1)
	t.Log(len(mySet))
}
