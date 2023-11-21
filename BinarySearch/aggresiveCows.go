package BinarySearch

import "sort"

// link - https://takeuforward.org/data-structure/aggressive-cows-detailed-solution/

func FindMaximumOfMinimumDistanceBetweenCows(stallsPos []int, size, cowsCount int) int {
	sort.Ints(stallsPos)
	low, high := 1, stallsPos[size-1]-stallsPos[0]
	res := 0
	for low <= high {
		mid := (low+high) >> 1
		if canWePlaceCowsWithMinDistance(stallsPos, size, cowsCount, mid) {
			res = mid
			low = mid+1
		} else {
			high = mid-1
		}
	}
	return res
}

func canWePlaceCowsWithMinDistance(stallsPos []int, size, cowsCount, minDistanceAllowed int) bool {
	cowsPlaced := 1
	cowPos := stallsPos[0]
	for _, pos := range stallsPos {
		if pos - cowPos >= minDistanceAllowed {
			cowsPlaced += 1
			cowPos = pos
			if cowsPlaced == cowsCount {
				return true
			}
		} else {
			continue
		}
	}
	return false
}