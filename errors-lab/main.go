package main

import (
	"errors"
	"fmt"
	"strconv"
)

// ============================================================
// 错误处理三件套：哨兵错误 / %w 包装 + errors.Is / 自定义类型 + errors.As
// ============================================================

// ---------- 题 1：哨兵错误（sentinel error）----------

var ErrNotFound = errors.New("ErrNotFound")

func findUser(id int) (string, error) {
	if id == 1 {
		return "harry", nil
	}
	return "", ErrNotFound
}

// ---------- 题 2：%w 包装 + errors.Is ----------

func loadProfile(id int) (string, error) {
	str, err := findUser(id)
	if err != nil {
		return "", fmt.Errorf("load profile %d: %w", id, err)
	}
	return str, nil
}

// ---------- 题 3：自定义错误类型 + errors.As ----------

type NotFoundError struct {
	ID int
}

func (e *NotFoundError) Error() string {
	return "user " + strconv.Itoa(e.ID) + " not found"
}

func findUser2(id int) (string, error) {
	if id == 1 {
		return "harry", nil
	}
	return "", &NotFoundError{ID: id}
}

func main() {
	// --- 题 1/2 验证 ---
	_, err := loadProfile(99)
	fmt.Println("err 长什么样:", err) // 期望能看到两层信息

	if errors.Is(err, ErrNotFound) {
		fmt.Print("same error 1")
	}

	if err == ErrNotFound {
		fmt.Print("same error 2")
	}

	// --- 题 3 验证 ---
	_, err2 := findUser2(99)
	var nfe *NotFoundError
	if errors.As(err2, &nfe) {
		fmt.Println("not found id:", nfe.ID)
	}
}
