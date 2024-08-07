package slice

// Reverse 将会完全创建一个新的切片，而不是直接在 src 上进行翻转。
func Reverse[T any](s []T) []T {
	var r = make([]T, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		r[i] = s[i]
	}
	return r
}

// ReverseSelf 會直接在 src 上进行翻转。
func ReverseSelf[T any](src []T) {
	for i, j := 0, len(src)-1; i < j; i, j = i+1, j-1 {
		src[i], src[j] = src[j], src[i]
	}
}
