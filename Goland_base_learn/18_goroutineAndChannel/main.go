package main

import (
	"fmt"
	"sync"
	"time"
)

type Lock interface {
	UseLock()
	GetUseCount() int
}

type SafeLock struct {
	mu       sync.Mutex
	useCount int
}

func (s *SafeLock) UseLock() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.useCount++
}

func (s *SafeLock) GetUseCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.useCount
}

type UnSafeLock struct {
	useCount int
}

func (u *UnSafeLock) UseLock() {
	u.useCount++
}

func (u *UnSafeLock) GetUseCount() int {
	return u.useCount
}

func OnlyReceive(ch <-chan int) {
	time.Sleep(4 * time.Second)
	for num := range ch {
		println("接收数据：", num)
	}

}

func OnlySend(ch chan<- int) {
	for i := 0; i < 5; i++ {
		println("发送数据前：", i)
		ch <- i
		println("发送数据后：", i)
		//time.Sleep(2 * time.Second)
	}
	println("通道即将关闭")
	close(ch)
}

func main() {
	/*channel 案例*/
	ch := make(chan int, 3)

	go OnlySend(ch)

	//OnlyReceive(ch)

	timeout := time.After(7 * time.Second)

	for {
		select {
		case v, ok := <-ch:
			if !ok {
				println("通道已关闭")
				return
			}
			fmt.Printf("主gocoutine接收到数据：%d\n", v)
			time.Sleep(1 * time.Second)
		case <-timeout:
			fmt.Println("操作超时")
			return
		default:
			fmt.Println("无数据，等待中")
		}
	}

}

/*goroutine 案例*/
//lock := &SafeLock{}
//
//for i := 0; i < 100; i++ {
//	go func() {
//		for i := 0; i < 10; i++ {
//			lock.UseLock()
//		}
//	}()
//}
//
//time.Sleep(time.Second)
//
//fmt.Printf("使用次数：%d\n", lock.GetUseCount())
