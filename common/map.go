package common

import "strconv"

func MapTo[T, U any](ts []T, f func(T) U) []U {
	res := make([]U, len(ts))
	for i, t := range ts {
		res[i] = f(t)
	}
	return res
}

func ToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
