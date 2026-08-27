package main

import (
	"errors"
	"fmt"
)

var _ = errors.New // 占位保住 import，用上 errors 后删掉这行

// ============================================================
// 错误处理三件套：哨兵错误 / %w 包装 + errors.Is / 自定义类型 + errors.As
// ============================================================

// ---------- 题 1：哨兵错误（sentinel error）----------

// TODO 1: 定义包级哨兵错误 ErrNotFound（用 errors.New）
// var ErrNotFound = ...

// TODO 2: findUser(id int) (string, error)
//   id == 1 → 返回 "harry", nil
//   其他    → 返回 "", ErrNotFound
func findUser(id int) (string, error) {
	return "", nil // 占位，替换掉
}

// ---------- 题 2：%w 包装 + errors.Is ----------

// TODO 3: loadProfile 调 findUser，出错时用 fmt.Errorf 包一层：
//   fmt.Errorf("load profile %d: %w", id, err)   ← %w 是关键，不是 %v
func loadProfile(id int) (string, error) {
	return "", nil // 占位，替换掉
}

// ---------- 题 3：自定义错误类型 + errors.As ----------

// TODO 5: 定义 NotFoundError struct，带一个 ID int 字段
//         给它实现 Error() string（指针 receiver），比如 "user 99 not found"

// TODO 6: findUser2(id int) (string, error)
//   id == 1 → "harry", nil
//   其他    → "", &NotFoundError{ID: id}
func findUser2(id int) (string, error) {
	return "", nil // 占位，替换掉
}

func main() {
	// --- 题 1/2 验证 ---
	_, err := loadProfile(99)
	fmt.Println("err 长什么样:", err) // 期望能看到两层信息

	// TODO 4: 用 errors.Is(err, ErrNotFound) 判断根因，打印结果
	//         再对比一行 err == ErrNotFound，先预测两者各是 true 还是 false，再跑

	// --- 题 3 验证 ---
	_, err2 := findUser2(99)
	// TODO 7: 用 errors.As 把 *NotFoundError 取出来，打印它的 ID 字段
	//   var nfe *NotFoundError
	//   if errors.As(err2, &nfe) { ... }
	_ = err2 // 写完 TODO 7 后删掉这行
}
