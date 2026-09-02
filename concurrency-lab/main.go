package main

import (
	"fmt"
	"time"
)

// ============================================================
// 并发第一课：goroutine + channel 三题
// 每题写完先预测，再 go run . 看；报 deadlock 就问"另一头在哪"
// ============================================================

// ---------- 题 1：main 退出，goroutine 全死 ----------

// 预测：main 里调它之后，这句话会不会打出来？
func fireAndForget() {
	go func() {
		fmt.Print("goroutine says hi\n")
	}()
	time.Sleep(10 * time.Millisecond)
}

// ---------- 题 2：用 channel 等一个 goroutine ----------

// go 一个匿名函数：打印 "working..."，time.Sleep(100 * time.Millisecond)，然后 done <- true
// 函数末尾 <-done 等它
// 预测：打印顺序？函数大概耗时多久？
func waitWithChannel() {
	done := make(chan bool)
	go func() {
		fmt.Print("working...")
		time.Sleep(100 * time.Millisecond)
		done <- true
	}()
	<-done
}

// ---------- 题 3：三个 goroutine 分段求和，channel 收回来 ----------

// 注意参数类型 chan<- int 表示"只发不收"的 channel，编译器帮你防止在里面误读
func sumRange(nums []int, out chan<- int) {
	var t int
	for _, n := range nums {
		t += n
	}
	out <- t
}

// 造一个 out channel，go 三次 sumRange，然后从 out 收三次相加返回
// 预测：三段的结果到达顺序固定吗？
func parallelSum(nums []int) int {
	out := make(chan int)
	n := 3
	go sumRange(nums[:n], out)
	go sumRange(nums[n:2*n], out)
	go sumRange(nums[2*n:], out)

	total := 0
	for i := 0; i < 3; i++ {
		total += <-out
	}

	return total
}

func main() {
	fmt.Println("--- 题 1 ---")
	fireAndForget()
	// 预测：无 ✓（函数里加了 Sleep 后才打出来）
	// 实验做过：Sleep 加在 fireAndForget 里

	fmt.Println("--- 题 2 ---")
	start := time.Now()
	waitWithChannel()
	fmt.Println("waited:", time.Since(start).Round(time.Millisecond)) // 预测：working... 然后 waited ≈100ms ✓

	fmt.Println("--- 题 3 ---")
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("sum =", parallelSum(nums)) // 期望 45
}
