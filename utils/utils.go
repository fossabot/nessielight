package utils

type any interface{}

func Filter[T any](slice []T, tester func(t T) bool) []T {
	var len int
	for _, v := range slice {
		if tester(v) {
			len++
		}
	}
	res := make([]T, 0, len)
	for _, v := range slice {
		if tester(v) {
			res = append(res, v)
		}
	}
	return res
}

func Reduce[T any, M any](s []T, f func(M, T) M, initValue M) M {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}
