```
package main

import (
	"fmt"
	"net"
	"sort"
	"sync"
)

var wg sync.WaitGroup
var Mutex sync.Mutex
var index int = 10

func one() {
	Mutex.Lock()
	index += 30
	Mutex.Unlock()
	wg.Done()
}

func two() {
	Mutex.Lock()
	index += 40
	Mutex.Unlock()
	wg.Done()
}

// 利用锁决绝资源抢占
func lock() {
	wg.Add(10)
	for i := 0; i < 5; i++ {
		go one()
		go two()
	}

	wg.Wait()
	fmt.Println(index)
	fmt.Println("主线程完事, 等待子线程")
}

func example() {
	//创建任务列表
	ports := make(chan int, 100)
	// 任务结果
	results := make(chan int)
	// 打开列表
	var openports []int
	//开启多线程等待任务
	for i := 0; i < cap(ports); i++ {
		go func(ports chan int, results chan int) {
			// 如果管道里面没有东西就阻塞
			for p := range ports {
				address := fmt.Sprintf("127.0.0.1:%d", p)
				conn, err := net.Dial("tcp", address)
				if err != nil {
					// 执行失败传0
					results <- 0
					continue
				}
				// 关闭请求
				conn.Close()
				// 成功传回p
				results <- p
			}
		}(ports, results)

	}
	go func() {
		// 派发任务
		for i := 1; i < 1024; i++ {
			// 这里放入任务的管道会被阻塞的线程抢占
			ports <- i
		}
	}()
	for i := 1; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf(" %d open\n", port)
	}
}

func main() {
	example()
}
```
