package Bot

import "reflect"

func find[T any](list []T, item T) int {
	for i := range list {
		if reflect.DeepEqual(item, (list)[i]) {
			return i
		}
	}
	return -1
}

func findComplex[T any, R any](list []T, item R, equal func(T, R) bool) int {
	for i := range list {
		if equal((list)[i], item) {
			return i
		}
	}
	return -1
}

func findWhere[T any](list []T, equal func(T) bool) int {
	for i := range list {
		if equal((list)[i]) {
			return i
		}
	}
	return -1
}

func removeItem[T any](list []T, item T) []T {
	idex := find(list, item)
	if idex != -1 {
		list[idex] = list[len(list)-1]
		return list[:len(list)-1]
	}
	return list
}

func removeComplex[T any, R any](list []T, item R, equal func(T, R) bool) []T {
	idex := findComplex(list, item, equal)
	if idex != -1 {
		list[idex] = list[len(list)-1]
		return list[:len(list)-1]
	}
	return list
}

func contains[T any](list []T, item T) bool {
	return find(list, item) != -1
}

func containsComplex[T any, R any](list []T, item R, equal func(T, R) bool) bool {
	for i := range list {
		if equal((list)[i], item) {
			return true
		}
	}
	return false
}
