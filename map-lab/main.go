package main

import (
	"fmt"
	"strings"
)

// ---------- 题 1：词频统计 ----------

// TODO 1: WordCount 统计每个单词出现次数
// 提示：strings.Fields(s) 把句子按空白切成 []string
// 注意：map 必须先 make，var 出来的 nil map 一写就 panic——可以先故意用 var 踩一次看报错
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, w := range strings.Fields(s) {
		m[w]++
	}
	return m
}

// ---------- 题 2：slice 去重 ----------

// TODO 2: dedup 去掉重复元素，保持首次出现的顺序
// 提示：map[int]bool 当 set 用，见过的跳过、没见过的记下并追加到结果
func dedup(nums []int) []int {
	var intArr []int
	var m = make(map[int]int)
	for _, n := range nums {
		m[n]++
		if m[n] == 1 {
			intArr = append(intArr, n)
		}
	}
	return intArr
}

func main() {
	fmt.Println(WordCount("go is fun and go is fast"))
	// 期望：map[and:1 fast:1 fun:1 go:2 is:2]（打印时 key 自动排序，遍历才是随机的）

	fmt.Println(dedup([]int{3, 1, 3, 2, 1, 3}))
	// 期望：[3 1 2]
}
