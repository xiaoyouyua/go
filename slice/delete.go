package main

import "fmt"

func Delete[T Number](idx int, arr []T) ([]T, T, error) {
	length := len(arr)
	if idx < 0 || idx >= length {
		var zero T
		return nil, zero, fmt.Errorf("下标超出范围，长度%d,下标%d", idx, len(arr))
	}
	res := arr[idx]
	for i := idx; i+1 < length; i++ {
		arr[i] = arr[i+1]
	}
	arr[length-1] = 0
	return arr, res, nil
}

type Number interface {
	~int | ~int64 | ~float64
}
