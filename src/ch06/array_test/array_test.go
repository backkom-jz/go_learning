package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var a [3]int

	for _, val := range a {
		t.Log(val)
	}

	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 4, 5}

	arr1[1] = 1
	t.Log(arr1, arr3)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	for idx, val := range arr3 {
		t.Log(idx, val)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	arr3_sec := arr3[2:]
	arr3_sec2 := arr3[:2]

	t.Log(arr3_sec, arr3_sec2)

}
