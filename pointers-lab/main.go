package main

import "fmt"

// ============================================================
// 指针三题：& 造链接 / * 开链接
// 阶段一 §6 补课
// ============================================================

// ---------- 题 1：swap ----------

// TODO 1: swap(a, b *int) 交换两个数（用 * 开链接）
func swap(a, b *int) {
	// 两三行
	t := *a
	*a = *b
	*b = t
}

// ---------- 题 2：值传 vs 指针传 ----------

type User struct {
	Name string
}

// TODO 2: renameByValue 收 User（复印件），把 Name 改成 "changed"
func renameByValue(u User) {
	// 一行
	u.Name = "changed"
}

// TODO 3: renameByPointer 收 *User（链接），把 Name 改成 "changed"
//
//	提示：p.Name 就行，Go 自动开链接，不用 (*p).Name
func renameByPointer(u *User) {
	// 一行
	u.Name = "changed"
}

// ---------- 题 3：值 receiver vs 指针 receiver ----------

type Counter struct {
	N int
}

// TODO 4: IncByValue —— 值 receiver，N++
// func (c Counter) ...

// TODO 5: IncByPointer —— 指针 receiver，N++
// func (c *Counter) ...

func main() {
	// --- 题 1 ---
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println("swap 后:", x, y) // 期望 2 1

	// --- 题 2：先预测两行各打印什么，再跑 ---
	u1 := User{Name: "harry"}
	renameByValue(u1)
	fmt.Println("值传后:", u1.Name) // 预测：____

	u2 := User{Name: "harry"}
	renameByPointer(&u2)
	fmt.Println("指针传后:", u2.Name) // 预测：____

	// --- 题 3：先预测，再跑 ---
	c := Counter{}
	// TODO 6: 调 IncByValue 三次、打印 c.N；再调 IncByPointer 三次、打印 c.N
	//         预测两次分别是几？
	_ = c // 写完 TODO 6 删掉这行
}
