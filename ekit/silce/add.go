package silce

import "src/ekit/internal/silce"

// Add 在index处添加元素
// src 是传入的切片
// element 插入的元素
// index 插入语的位置
// index 范围应为[0, len(src)]
// 如果index == len(src) 则表示往末尾添加元素
func Add[Src any](src []Src, element Src, index int) ([]Src, error) {
	res, err := silce.Add[Src](src, element, index)
	return res, err
}
