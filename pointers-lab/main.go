package main

import "fmt"

// ============================================================
// 指针三题：& 造链接 / * 开链接
// 阶段一 §6 补课
// ============================================================

// ---------- 题 1：swap ----------

func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

// ---------- 题 2：值传 vs 指针传 ----------

type User struct {
	Name string
}

func renameByValue(u User) {
	u.Name = "changed"
}

// 提示：p.Name 就行，Go 自动开链接，不用 (*p).Name
func renameByPointer(u *User) {
	u.Name = "changed"
}

// ---------- 题 3：值 receiver vs 指针 receiver ----------

type Counter struct {
	N int
}

func (c Counter) IncByValue() {
	c.N++
}

func (c *Counter) IncByPointer() {
	c.N++
}

func main() {
	// --- 题 1 ---
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println("swap 后:", x, y) // 期望 2 1

	// --- 题 2：先预测两行各打印什么，再跑 ---
	u1 := User{Name: "harry"}
	renameByValue(u1)
	fmt.Println("值传后:", u1.Name) // 预测：harry ✓

	u2 := User{Name: "harry"}
	renameByPointer(&u2)
	fmt.Println("指针传后:", u2.Name) // 预测：changed ✓

	// --- 题 3：先预测，再跑 ---

	c := Counter{}
	c.IncByValue()
	c.IncByValue()
	c.IncByValue()
	fmt.Printf("value: %d \n", c.N) // 0

	c.IncByPointer()
	c.IncByPointer()
	c.IncByPointer()
	fmt.Printf("value: %d \n", c.N) // 3
}
