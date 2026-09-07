package main

import (
	"fmt"
	"sync"
	"time"
)

// ============================================================
// 并发第二课：WaitGroup + ping-pong（正式 7 题的前 2 题）
// 老规矩：写完先在注释里写预测，再 go run .；deadlock 就问"另一头在哪"
// ============================================================

// ---------- 题 1：3 个 goroutine 打印，WaitGroup 等它们全部结束 ----------

// 要做的事：
//   1. var wg sync.WaitGroup
//   2. for i := 1; i <= 3; i++ 循环里：wg.Add(1)，然后 go worker(i, &wg)
//   3. 循环外 wg.Wait()
//   4. 最后打印 "all workers done"
//
// worker 自己：第一行 defer wg.Done()，然后 Sleep(随机一点，比如 time.Duration(id*30)*time.Millisecond)，打印 "worker <id> done"
//
// 预测 A：三行 "worker x done" 的顺序固定吗？
// 预测 B：把 wg.Wait() 注释掉再跑，会打出几行？
// 预测 C：worker 参数改成 wg sync.WaitGroup（不带 *）会怎样？（先想再试，试完改回来）
func worker(id int, wg *sync.WaitGroup) {
	// TODO
}

func waitGroupDemo() {
	// TODO
}

// ---------- 题 2：无缓冲 channel 做 ping-pong ----------
// 等题 1 跑通、预测对完，再来看这题，讲解在对话里。

// 要做的事：
//   两个 channel：ping := make(chan int)、pong := make(chan int)
//   pinger 是一个 goroutine：循环 rounds 次 { n := <-ping; 打印 "ping", n; pong <- n + 1 }
//   main（pingPong 函数本体）当 pong 选手：先发球 ping <- 0，然后循环 rounds 次 { n := <-pong; 打印 "pong", n; 如果不是最后一轮就 ping <- n + 1 }
//
// 预测 A：输出是不是严格 ping/pong 交替？为什么能保证？
// 预测 B：pong 选手多循环一次（rounds+1）会发生什么？报错信息第一行是什么？
func pinger(ping <-chan int, pong chan<- int, rounds int) {
	// TODO
}

func pingPong(rounds int) {
	// TODO
}

func lesson2() {
	fmt.Println("--- 第二课 题 1: WaitGroup ---")
	waitGroupDemo()

	fmt.Println("--- 第二课 题 2: ping-pong ---")
	pingPong(3)
	_ = time.Millisecond // 写完 worker 用到 time 后删掉这行
}
