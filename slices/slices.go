package slices

import "fmt"

func Map[T, V any](col []T, fn func(T) V) []V {
	r := make([]V, len(col))
	for i, e := range col {
		r[i] = fn(e)
	}
	return r
}

func Filter[T any](col []T, pred func(T) bool) []T {
	r := []T{}
	for _, e := range col {
		if pred(e) {
			r = append(r, e)
		}
	}
	return r
}

func Contains[T any](col []T, pred func(T) bool) bool {
	for _, e := range col {
		if pred(e) {
			return true
		}
	}
	return false
}

func Merge[T1, T2, V any](col1 []T1, col2 []T2, merge func(T1, T2) V) ([]V, error) {
	if len(col1) != len(col2) {
		return nil, fmt.Errorf("collections must be of same length")
	}
	r := make([]V, len(col1))
	for i := 0; i < len(col1); i++ {
		r = append(r, merge(col1[i], col2[i]))
	}
	return r, nil
}

func GroupBy[T any, K comparable](col []T, groupBy func(item T) K) map[K][]T {
	result := map[K][]T{}
	for _, item := range col {
		key := groupBy(item)
		result[key] = append(result[key], item)
	}
	return result
}
