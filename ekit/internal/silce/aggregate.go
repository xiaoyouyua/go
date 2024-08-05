package silce

import (
	"src/ekit/silce"
)

// MAX
// 该方法会假设你至少会传入一个值
// 在使用 float32 或者 float64 的时候要小心精度问题
func Max[T silce.RealNumber](ts []T) T {
	max := ts[0]
	for i := 1; i < len(ts); i++ {
		if max < ts[i] {
			max = ts[i]
		}
	}
	return max
}

// MIN
// 该方法会假设你至少会传入一个值
// 在使用 float32 或者 float64 的时候要小心精度问题
func Min[T silce.RealNumber](ts []T) T {
	min := ts[0]
	for i := 1; i < len(ts); i++ {
		if min > ts[i] {
			min = ts[i]
		}
	}
	return min
}

// Sum 求和
// 在使用 float32 或者 float64 的时候要小心精度问题
func Sum[T silce.RealNumber](ts []T) T {
	var sum T
	for _, t := range ts {
		sum += t
	}
	return sum
}
