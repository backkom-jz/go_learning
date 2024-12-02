package pointer

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	// 1.定义一个普通变量
	var num = 666

	fmt.Println(num)
	// 2.定义一个指针变量
	var num2 *int = &num
	fmt.Printf("%p\n", &num)
	fmt.Printf("%p\n", num2)
	fmt.Printf("%T\n", num2)
	// 3.通过指针变量操作指向的存储空间
	*num2 = 888

	// 4.指针变量操作的就是指向变量的存储空间
	t.Log(num, *num2)

}
