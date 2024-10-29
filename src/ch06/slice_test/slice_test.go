package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var slice0 []int
	t.Log(len(slice0), cap(slice0))

	slice0 = append(slice0, 1)
	t.Log(len(slice0), cap(slice0))

	slice1 := []int{1, 2, 3, 4, 5}
	t.Log(len(slice1), cap(slice1))

	slice2 := make([]int, 3, 5) // capacity 容量
	t.Log(len(slice2), cap(slice2))

	slice2 = append(slice2, 1)
	t.Log(len(slice2), cap(slice2))
}

// 共享的存储结构

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{
		"January", "February", "March", "April", "May", "Jun",
	}
	Q1 := year[0:3]
	Q2 := year[0:5]
	t.Log(Q1, len(Q1), cap(Q1))
	t.Log(Q2, len(Q2), cap(Q2))
}
