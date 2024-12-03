package concurrent

import (
	"fmt"
	"testing"
	"time"
)

// 无缓冲channel适用于 通知 B要第一时间知道A是否已完成
// 有缓存channel适用于消费者和生产者的通信
/**
go中channel应用场景
	1、消息传递、消息过滤
	2、信号广播
	3、事件订阅和广播
	4、任务分发
	5、结果汇总
	6、并发控制
	7、同步和异步
*/
func TestChannel(t *testing.T) {
	// 不要通过共享内存来进行通信，而要通过通信来实现内存共享

	// php python java 多线程编程的时候 两个goroutine之间的通信最常用的方式是一个全局
	// 也会 提供消息队列的机制 python-queue java 消费者和生产者之间的关系
	// channel 再加上语法糖让使用channel更加简单

	var msg chan string
	if msg == nil {
		fmt.Println(msg)
	}
	msg = make(chan string, 1) // channel 初始化为0的时候，放值的时候会阻塞

	go func(msg chan string) { // go有一种happen-before机制，可以保障
		data := <-msg
		fmt.Println(data)
	}(msg)
	msg <- "bobby" //放值到channel中
	data := <-msg

	// 有缓冲 无缓冲
	fmt.Println(data)

	// waitgroup 如果少了done的调用 无缓冲的channel也容易出现deadlock
}

func TestChannelForRange(t *testing.T) {
	var msg chan int
	if msg == nil {
		fmt.Println(msg)
	}
	msg = make(chan int, 2) // channel 初始化为0的时候，放值的时候会阻塞

	go func(msg chan int) { // go有一种happen-before机制，可以保障
		for data := range msg {
			fmt.Println(data)
		}
		fmt.Println("all done")
	}(msg)
	msg <- 1 //放值到channel中
	msg <- 2 //放值到channel中

	//data := <-msg
	close(msg) // 已经关闭的channel 只能取值 不能再放值

	//msg <- 3
	time.Sleep(time.Second * 10)
}

func producer(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i * i
	}
	close(out)
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Printf("num=%d\r\n", num)
	}
}

func TestSignChannel(t *testing.T) {
	// 默认情况下 channel是双向的
	// channel作为参数进行传递，希望对方是单向的

	//var ch1 chan int       // 双向channel
	//var ch2 chan<- float64 // 单向channel，只能写入float64的数据
	//var ch3 <-chan int     // 单向的 只能读取

	//c := make(chan int, 3)
	//var send chan<- int = c // send-only
	//var read <-chan int = c // read-only
	//
	//send <- 1
	//<-read

	c := make(chan int)
	go producer(c)
	go consumer(c)

	time.Sleep(10 * time.Second)
}

var number, letter = make(chan bool), make(chan bool)

func printNumber() {
	i := 1
	for {
		<-number
		fmt.Printf("%d%d", i, i+1)
		i += 2
		letter <- true
	}
}
func printLetter() {
	i := 0
	str := "ABCDEFGHIJKLMNOPQR"
	for {
		<-letter
		if i >= len(str) {
			return
		}
		fmt.Print(str[i : i+2])
		i += 2
		number <- true
	}
}

// 常见channel面试题
func TestPrintChannel(t *testing.T) {
	/**
	使用两个goroutine交替打印序列，一个goroutine打印字母，一个goroutine打印数组
	12AB34CD56EF...
	*/
	go printNumber()
	go printLetter()
	number <- true

	time.Sleep(10 * time.Second)
}
