package main

import (
	"fmt"
)

// ---------- 题 1：隐式实现 + 多态 ----------

// TODO 1: 定义 Shape 接口，只有一个方法 Area() float64
type Shape interface {
	// 一行
	Area() float64
}

type Circle struct {
	R float64
}

type Rectangle struct {
	W, H float64
}

// TODO 2: 给 Circle 写 Area()（值 receiver），公式 3.14159 * R * R
func (c Circle) Area() float64 {
	return 3.14159 * c.R * c.R
}

// TODO 3: 给 Rectangle 写 Area()（值 receiver）
func (r Rectangle) Area() float64 {
	return r.W * r.H
}

// TODO 4: totalArea 接收 []Shape，返回总面积
func totalArea(shapes []Shape) float64 {
	var total float64 = 0
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}

// ---------- 题 2：类型 switch ----------

// TODO 5: describe 用 type switch 区分 int / string / 其他，各打印一句
func describe(x any) {
	switch v := x.(type) {
	case int:
		fmt.Println("int: ", v)
	case string:
		fmt.Println("string: ", v)
	default:
		fmt.Println("other: ", v)
	}
}

func main() {
	runReference()

	shapes := []Shape{Circle{R: 1}, Rectangle{W: 2, H: 3}}
	fmt.Printf("total=%.2f\n", totalArea(shapes)) // 期望 9.14

	describe(42)   // 期望：int 42
	describe("hi") // 期望：string hi
	describe(3.14) // 期望：其他类型
}
