package KadanesAlgo

import "math"

func MaxSubArray(nums []int) int {
	maxSoFar, maxEndingHere := math.MinInt64, 0
	for _, num := range nums {
		maxEndingHere += num 
		// update maxSoFar
		if maxEndingHere > maxSoFar {
			maxSoFar = maxEndingHere
		}
		// reset maxEndingHere if sum till then is negative 
		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
	}
	return maxSoFar
}