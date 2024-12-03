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

func TestArrayPointer(t *testing.T) {
	var arr [3]int = [3]int{1, 3, 5}
	var p *[3]int
	p = &arr
	fmt.Printf("%p\n", &arr) // 0xc0420620a0
	fmt.Printf("%p\n", p)    // 0xc0420620a0
	fmt.Println(&arr)        // &[1 3 5]
	fmt.Println(p)           // &[1 3 5]

	// 指针指向数组之后操作数组的几种方式
	//1.直接通过数组名操作
	arr[1] = 6
	fmt.Println(arr[1])
	// 2.通过指针间接操作
	(*p)[1] = 7
	fmt.Println((*p)[1])
	fmt.Println(arr[1])
	//
	//3.通过指针间接操作 p[1] = 8 fmt.Println(p[1]) fmt.Println(arr[1])

	// 注意点: Go语言中的指针, 不支持+1 -1和++ --操作 *(p + 1) = 9 // 报错
	//fmt.Println(*p++) // 报错
	fmt.Println(arr[1])

}

func TestPointerSlice(t *testing.T) {
	// 1.定义一个切片
	var sce []int = []int{1, 3, 5}
	// 2.打印切片的地址
	// 切片变量中保存的地址, 也就是指向的那个数组的地址 sce = 0xc0420620a0 fmt.Printf("sce = %p\n",sce )
	fmt.Println(sce) // [1 3 5]
	// 切片变量自己的地址, &sce = 0xc04205e3e0
	fmt.Printf("&sce = %p\n", &sce)
	fmt.Println(&sce) // &[1 3 5]
	// 3.定义一个指向切片的指针
	var p *[]int
	// 因为必须类型一致才能赋值, 所以将切片变量自己的地址给了指针
	p = &sce
	// 4.打印指针保存的地址
	// 直接打印p打印出来的是保存的切片变量的地址 p = 0xc04205e3e0 fmt.Printf("p = %p\n", p)
	fmt.Println(p) // &[1 3 5]
	// 打印*p打印出来的是切片变量保存的地址, 也就是数组的地址 *p = 0xc0420620a0 fmt.Printf("*p = %p\n", *p)
	fmt.Println(*p) // [1 3 5]
	// 5.修改切片的值
	// 通过*p找到切片变量指向的存储空间(数组), 然后修改数组中保存的数据 (*p)[1] = 666
	fmt.Println(sce[1])
}

func TestPointerMap(t *testing.T) {

	var dict map[string]string = map[string]string{"name": "lnj", "age": "33"}
	var p *map[string]string = &dict
	(*p)["name"] = "zs"
	fmt.Println(dict)
}

type Student struct {
	name string
	age  int
}

func TestPointerStruct(t *testing.T) {
	// 创建时利用取地址符号获取结构体变量地址
	var p1 = &Student{"lnj", 33}
	fmt.Println(p1) // &{lnj 33}
	// 通过new内置函数传入数据类型创建
	// 内部会创建一个空的结构体变量, 然后返回这个结构体变量的地址
	var p2 = new(Student)
	fmt.Println(p2) // &{ 0}
}
