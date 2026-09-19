package main

import (
	"fmt"
	"time"
)

// ============================================================
// 并发第四课：select + 超时（正式 7 题的第 4 题）
// 老规矩：先在注释里写预测 → go run . → 再对答案
// ============================================================

// ---------- 题 4a：一个会慢的工人 ----------
//
// slowWork：睡 d 那么久，然后把结果写进 out。
//
// 要做的事：
//  1. time.Sleep(d)
//  2. out <- 42
func slowWork(d time.Duration, out chan<- int) {
	time.Sleep(d)
	out <- 42
}

// ---------- 题 4b：select 做超时 ----------
//
// fetchWithTimeout：等 slowWork 的结果，最多等 timeout 这么久。
//   - 拿到结果 → 返回 (结果, nil)
//   - 超时     → 返回 (0, error)
//
// 要做的事：
//  1. out := make(chan int, 1)   ← 为什么给 1 格缓冲？见文末思考题
//  2. go slowWork(work, out)
//  3. select {
//     case r := <-out:                 return r, nil
//     case <-time.After(timeout):      return 0, fmt.Errorf("超时：等了 %v", timeout)
//     }
//
// 预测（写完再跑，填在下面）：
//   - fetchWithTimeout(work=500ms, timeout=2s) → 结果是？ 42,nil
//   - fetchWithTimeout(work=3s,   timeout=1s) → 结果是？大约几秒后返回？ 0, error; 1s
//   - 第二种情况下，slowWork 那个 goroutine 怎么样了？
func fetchWithTimeout(work, timeout time.Duration) (int, error) {
	out := make(chan int, 1)
	go slowWork(work, out)
	select {
	case r := <-out:
		return r, nil
	case <-time.After(timeout):
		return 0, fmt.Errorf("超时：等了 %v", timeout)
	}
}

func lesson4() {
	fmt.Println("--- 第四课 题 4: select 超时 ---")

	start := time.Now()
	r, err := fetchWithTimeout(500*time.Millisecond, 2*time.Second)
	fmt.Printf("快活: r=%d err=%v 用时=%v\n", r, err, time.Since(start).Round(time.Millisecond))

	start = time.Now()
	r, err = fetchWithTimeout(3*time.Second, 1*time.Second)
	fmt.Printf("慢活: r=%d err=%v 用时=%v\n", r, err, time.Since(start).Round(time.Millisecond))
}

// 思考题（跑通后再答，不用写代码）：
//  Q: 第 1 步如果写成 make(chan int) 无缓冲，超时那次会发生什么？不会
