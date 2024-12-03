package concurrent

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg1 sync.WaitGroup

// 我们新的需求 我可以主动退出监控程序
var stop = make(chan struct{})

func cpuInfo(ctx context.Context) {
	//  这里能拿到一个请求的ID
	fmt.Printf("tranceid:%s\r\n", ctx.Value("tranceid"))
	// 记录一些日志，这次请求的ID
	defer wg1.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("退出监控")
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("cpu 的信息")
		}
	}
}
func TestContext(t *testing.T) {

	// 渐进的方式
	// 有一个goroutine监控cpu的信息
	wg1.Add(1)
	//  context 包提供了三种函数 WithCancel WithTimeout WithValue
	// 如果你的goroutine 函数中，如果希望被控制，超时、传值，但是我不希望影响到我原来的接口信息的时候
	// 函数参数中的第一个参数尽量的要加上一个ctx
	//ctx, cancel := context.WithCancel(context.Background())
	//ctx2, _ := context.WithCancel(ctx)

	// 2. timeout 主动 超时
	ctx, _ := context.WithTimeout(context.Background(), 6*time.Second)

	// 3 withDeadline 在时间点cancel

	// 4 withValue
	valueCtx := context.WithValue(ctx, "tranceid", "123")
	go cpuInfo(valueCtx)

	time.Sleep(6 * time.Second)
	//cancel()

	wg1.Wait()
	fmt.Println("监控完成")
}
