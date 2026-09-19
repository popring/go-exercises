package main

import "fmt"

// ============================================================
// 并发第三课：Worker Pool（正式 7 题的第 3 题）
// 老规矩：先在注释里写预测 → go run . → 再对答案
// ============================================================

// ---------- 题 3：5 个 worker 从 jobs channel 取任务，结果写 results channel ----------

// poolWorker：一个工人。不停从 jobs 里取任务，算完写进 results。
//
// 要做的事：
//  1. for j := range jobs {  ← 一直取，直到 jobs 被 close
//  2. 循环体里：打印一句 "worker %d 处理 job %d\n"（看得见谁干了哪件活）
//  3. 把结果写进 results：results <- j * j
//  4. 循环结束后打印 "worker %d 下班\n"
//
// 提示：参数里的 <-chan / chan<- 箭头方向 = 这个函数只能读 / 只能写
func poolWorker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker %d 处理 job %d\n", id, j)
		results <- j * j
	}
	fmt.Printf("worker %d 下班\n", id)
}

// workerPool：老板。派活 + 收货。
//
// 要做的事：
//  1. jobs := make(chan int, 9)    ← 带缓冲，老板发完就走不用等
//  2. results := make(chan int, 9)
//  3. 起 5 个工人：for w := 1; w <= 5; w++ { go poolWorker(w, jobs, results) }
//  4. 发 9 个任务：for j := 1; j <= 9; j++ { jobs <- j }
//  5. close(jobs)                  ← 关键：不关，工人的 for range 永远等下一个，程序卡死
//  6. 收 9 个结果：for i := 0; i < 9; i++ { total += <-results }
//  7. 返回 total
//
// 预测（写完再跑）：
//   - 9 个 job 会不会按 1..9 的顺序打印？ 不会
//   - total 应该是多少？（1²+2²+...+9²）
func workerPool() int {
	jobs := make(chan int, 9)
	results := make(chan int, 9)
	total := 0

	for w := 1; w <= 5; w++ {
		go poolWorker(w, jobs, results)
	}
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	for range 9 {
		total += <-results
	}

	return total
}

func lesson3() {
	fmt.Println("--- 第三课 题 3: worker pool ---")
	fmt.Println("total =", workerPool())
}
