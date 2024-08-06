package biz

import "fmt"

// 要点1 函数的组成：func关键词、函数名、入参、出参、函数体{}，判断函数的唯一性：函数名+入参+出参
// 要点2 函数的可见范围，由函数名的首字母大小写决定的

// 初始化函数，当package被引用时，init函数会被自动调用
func init() {
	fmt.Println("biz.init")
}

// Add return sum a and b.
func Add(a, b int) int {
	c := a + b
	return c
}
