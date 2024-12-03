package concurrent

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var done bool
var lock1 sync.Mutex

// channel 多线程安全
var done2 = make(chan struct{})

// 很多时候我并不会多个goroutine写同一个channel

func g1() {
	time.Sleep(time.Second)
	lock1.Lock()
	defer lock1.Unlock()
	done = true
}

func g2() {
	fmt.Println(time.Second * 2)
	lock1.Lock()
	defer lock1.Unlock()
	done = true
}

func TestSelect(t *testing.T) {
	// select 类似与switch case语句 但是select的功能和我们操作linux里面提供的io的select、poll、epoll
	// select 主要作用于多个channel

	// 现在有个需求，我们现在有两个goroutine都在执行，但是我们在住的goroutine中，当一个执行完成一个，这个时候我会立马知道
	go g1()
	go g2()
	for {
		if done {
			fmt.Println("done")
			time.Sleep(10 * time.Millisecond)
			return
		}

	}
}

func gg1(ch chan struct{}) {
	time.Sleep(time.Second)
	ch <- struct{}{}
}

func gg2(ch chan struct{}) {
	time.Sleep(2 * time.Second)
	ch <- struct{}{}
}

func TestSelectChannel(t *testing.T) {
	g1Channel := make(chan struct{}, 1)
	g2Channel := make(chan struct{}, 2)
	//g1Channel <- struct{}{}
	//g2Channel <- struct{}{}
	go gg1(g1Channel)
	go gg2(g2Channel)

	// 我要监控多个channel 任何一个channel都能知道
	// 1、某一个分支就绪了就执行该分支 2 如果两个都就绪了，先执行那个，随机的,目的是什么，防止饥饿
	//for {
	//	select {
	//	case <-g1Channel:
	//		fmt.Println("g1 done")
	//	case <-g2Channel:
	//		fmt.Println("g2 done")
	//	default:
	//		time.Sleep(10 * time.Second)
	//		fmt.Println("default")
	//	}
	//}

	timer := time.NewTimer(time.Second)
	for {
		select {
		case <-g1Channel:
			fmt.Println("g1 done")
		case <-g2Channel:
			fmt.Println("g2 done")
		case <-timer.C:
			fmt.Println("over ")
			return
		default:
			time.Sleep(10 * time.Second)
			fmt.Println("default")
		}
	}

	fmt.Println("done2")
}
