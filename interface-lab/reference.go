package main

import "fmt"

// ============================================================
// 参考样例：照这个形状去填 main.go 里的 TODO
// 这里用 Speaker/Dog/Cat，跟你要写的 Shape/Circle/Rectangle 结构完全一样
// ============================================================

// 1) 接口：只列方法签名，不写实现，不写 implements
type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

// 2) 方法：值 receiver。写完这两个方法，Dog / Cat 自动就是 Speaker 了
func (d Dog) Speak() string {
	return d.Name + " says woof"
}

func (c Cat) Speak() string {
	return c.Name + " says meow"
}

// 3) 消费接口的函数：参数是接口类型，不关心具体是 Dog 还是 Cat
func speakAll(speakers []Speaker) {
	for _, s := range speakers {
		fmt.Println("  ", s.Speak())
	}
}

// 4) type switch：打开 any 的正规方式
func kindOf(x any) {
	switch v := x.(type) {
	case bool:
		fmt.Println("   bool", v)
	case []int:
		fmt.Println("   []int, 长度", len(v))
	default:
		fmt.Println("   不认识的类型")
	}
}

func runReference() {
	fmt.Println("== 参考样例 ==")
	speakAll([]Speaker{Dog{Name: "旺财"}, Cat{Name: "咪咪"}})
	kindOf(true)
	kindOf([]int{1, 2, 3})
	kindOf("字符串没进 case")
	fmt.Println("== 下面是你的作业 ==")
}
