package main

import (
	"fmt"
	"os"
)

// ============================================================
// defer / panic / recover 三题
// ============================================================

// ---------- 题 1：多个 defer 的执行顺序 ----------

// 先在注释里预测打印顺序，再跑
func deferOrder() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("done")
}

// ---------- 题 2：defer f.Close() ----------

// 然后 f.WriteString(content)，返回 err
// 规则：defer Close 紧跟在 err 检查之后，中间不插别的逻辑
func writeNote(path, content string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err2 := f.WriteString(content)
	return err2
}

// ---------- 题 3：recover 兜住 panic ----------

// 函数开头写一个 defer 的匿名函数，里面 recover()，
// 如果 recover 返回值不是 nil，就把 err 设成 fmt.Errorf("recovered: %v", r)
// 函数体直接写 result = a / b（b == 0 时 Go 自己会 panic，不用手动判）
// 提示：defer 里改 err 能生效，靠的是命名返回值——你在 todo-cli 已经用过
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()
	result = a / b
	return
}

func main() {
	// --- 题 1 ---
	deferOrder() // 预测：done 3 2 1 ✓

	// --- 题 2 ---
	if err := writeNote("note.txt", "hello defer\n"); err != nil {
		fmt.Println("write failed:", err)
	}
	data, _ := os.ReadFile("note.txt")
	fmt.Print("file content: ", string(data))
	os.Remove("note.txt")

	// --- 题 3：先预测两行各打印什么 ---
	r1, err1 := safeDivide(10, 2)
	fmt.Println("10/2 =", r1, err1) // 预测：5, nil ✓

	r2, err2 := safeDivide(10, 0)
	fmt.Println("10/0 =", r2, err2) // 预测：0, recovered ✓
}
