package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))
	//s[1] = '3' //string是不可变的byte slice
	//s = "\xE4\xB8\xA5" //可以存储任何二进制数据
	s = "\xE4\xBA\xBB\xFF"
	t.Log(s)
	t.Log(len(s))
	s = "中"
	t.Log(len(s)) //是byte数

	c := []rune(s)
	t.Log(len(c))
	//	t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
	}
}

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")

	for _, part := range parts {
		t.Log(part)
	}

	t.Log(strings.Join(parts, "-"))

}

func TestStringConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)

	s2, ok := strconv.Atoi("10")
	if ok == nil {
		t.Log(10 + s2)
	}

}
