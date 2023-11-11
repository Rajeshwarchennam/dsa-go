package MinMaxHelpers

import "golang.org/x/exp/constraints"

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInList[T constraints.Ordered](list []T) T {
	if len(list) == 0 {
		return *new(T)
	}
	max := list[0]
	for _, element := range list {
		if element > max {
			max = element
		}
	}
	return max
}

func MinInList[T constraints.Ordered](list []T) T {
	if len(list) == 0 {
		return *new(T)
	}
	min := list[0]
	for _, element := range list {
		if element < min {
			min = element
		}
	}
	return min
}
