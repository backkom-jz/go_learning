package concurrent

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//  python java php 多线程编程 多进程编程
// 内存、线程切换  web2.0 用户级线程【协程】
// asyncio-python php- swoole java -netty
// 内存占用小（2k） 切换快
// go语言的协程 goroutine 只有协程可用 非常方便

func asyncPrint() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("bobby")
	}

}

// 主协程
func TestGoroutine(t *testing.T) {
	// 主死随从
	//go asyncPrint()

	//go func() {
	//	for {
	//		time.Sleep(2 * time.Second)
	//		fmt.Println("bobby")
	//	}
	//}()

	// 1 闭包 2 for循环问题 每个变量会复用
	//for i := 0; i < 100; i++ {
	//	go func() {
	//		fmt.Println(i)
	//	}()
	//}

	for i := 0; i < 100; i++ {
		// 第一种解决方案
		//tmp := i
		//go func() {
		//	fmt.Println(tmp)
		//}()

		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	fmt.Println("main goroutine")
	time.Sleep(10 * time.Second)
}

// GMP
// G goroutine
// M processor [调度器、队列] 【映射】
// P thread-线程池【linux】 数量由 runtime.GOMAXPROCS() 函数控制，默认等于机器上的 CPU 核心数
// 调度过程：
//当你创建一个 goroutine 时，会创建一个新的 G，并将其放入 P 的 G 队列。
//P 会从队列中取出一个 G，并将其分配给 M，然后 M 执行这个 G。
//如果 M 完成了 G 的执行，P 会继续从队列中取下一个 G 让 M 执行。
//如果 P 的本地队列为空，P 会从全局队列或其他 P 的队列中偷取 G 来执行。
//GMP 优点：
//高效调度：goroutine 是非常轻量级的，Go 的调度器能够快速地将 goroutine 映射到系统线程中，提高并发执行的效率。
//自动伸缩：Go 运行时调度器能够根据需要动态调整 M（操作系统线程）的数量，以适应负载变化。
//简化并发编程：开发者只需要创建 goroutine，无需关心底层线程的管理，Go 运行时自动负责调度。

func TestWaitGroup(t *testing.T) {

	// goroutine  如何通知主的goroutine自己结束了，主的goroutine如何知道子的goroutine已经结束了
	var wg sync.WaitGroup

	// 我要监控多少个goroutine执行结束
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)

		}(i)
	}
	// 等到监控结束
	wg.Wait()
	fmt.Println("all done")

	// wait group 主要用于goroutine的执行等待，ADD方法和DONE方法配套
}

// 互斥锁  资源竞争
var total int
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock()
		total += 1
		lock.Unlock()
	}

}

// 锁能复制 复制后就会失去锁的作用
func sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock()
		total -= 1
		lock.Unlock()
	}
}
func TestLock(t *testing.T) {

	wg.Add(2)

	go add()
	go sub()

	wg.Wait()
	fmt.Println(total)
}

// atomic 原子包 数字的加减

// 读写锁
// 锁的本子是将并行的代码串行化，使用lock肯定会影响性能
// 即使是设计锁，也尽可能保证并行
// 我们有两组协程 其中一组负责写数据，另一组负责读数据，web系统绝大多数读多写少
// 虽有有多个goroutine 但是仔细分析我们会发现 读协程之间应该并发，读与写之间应该串行，读与读之间也不应该并行

func TestRwLock(t *testing.T) {
	var num int
	var rwlock sync.RWMutex
	var wg sync.WaitGroup
	wg.Add(6)
	go func() {
		defer wg.Done()
		rwlock.Lock()
		defer rwlock.Unlock()
		time.Sleep(5 * time.Second)
		num = 12
		fmt.Println("write lock")

	}()

	time.Sleep(time.Second)
	// 读的goruntine
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for {
				rwlock.RLock()
				time.Sleep(500 * time.Millisecond)
				fmt.Printf("read lock %d", num)
				rwlock.RUnlock()
			}

		}()
	}

	wg.Wait()
}
