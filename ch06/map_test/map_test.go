package map_test

import "testing"

func TestMapInit(t *testing.T) {
	m := make(map[int]int)
	m[1] = 1
	m[2] = 2
	t.Log(m)

	m1 := map[int]int{1: 1, 2: 2, 3: 9}
	t.Log(m1[2])

	m3 := make(map[int]int, 10)
	t.Logf("len m4=%d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])

	m1[2] = 0
	t.Log(m1[2])

	if v, ok := m1[3]; ok {
		t.Log(v)
	} else {
		t.Log("key 3 not found")
	}

}

func TestTravelMap(t *testing.T) {
	map1 := map[int]int{1: 1, 2: 2, 3: 9}
	for k, v := range map1 {
		t.Log(k, v)
	}
}
