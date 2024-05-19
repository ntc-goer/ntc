package ntc

import "reflect"

func RemoveItem[T any](its []T, i T) []T {
	res := make([]T, 0)
	for _, it := range its {
		if !reflect.DeepEqual(it, i) {
			res = append(res, i)
		}
	}
	return res
}

func Filter[T any](its []T, f func(i T) bool) []T {
	res := make([]T, 0)
	for _, it := range its {
		if f(it) {
			res = append(res, it)
		}
	}
	return res
}

func Contain[T any](its []T, i T) bool {
	for _, it := range its {
		if reflect.DeepEqual(it, i) {
			return true
		}
	}
	return false
}
