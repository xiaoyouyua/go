package slice

// 差集 ，只支持comparable 类型
// 以去重
// 并且返回值的顺序是不确定的
func DiffSet[T comparable](src, dst []T) []T {
	srcMap := toMap[T](src)
	for _, v := range dst {
		delete(srcMap, v)
	}

	var ret = make([]T, 0, len(srcMap))
	for v := range srcMap {
		ret = append(ret, v)
	}
	return ret
}

// DiffSetFunc 差集，已去重
// 你应该优先使用 DiffSet
func DiffSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	var ret = make([]T, 0, len(src))
	for _, v := range src {
		if !ContainsFunc[T](dst, func(src T) bool { return equal(src, v) }) {
			ret = append(ret, v)
		}
	}
	return deduplicateFunc[T](ret, equal)
}
