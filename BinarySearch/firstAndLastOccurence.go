package BinarySearch

//link - https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array
func FindFirstAndLastOccurence(nums []int, target int) []int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low+high) >> 1
		if nums[mid] == target {
			return []int{findFirstOccurence(nums, 0, mid, target), findLastOccurence(nums, mid, len(nums)-1, target)}
		}
		if nums[mid] > target {
			high = mid-1
		} else {
			low = mid+1
		}
	}

	return []int{-1,-1}
}


// finds first occurence of target in non-decreasing array where target is highest elemt 
func findFirstOccurence(nums []int, start, end, target int) int {
	for start <= end {
		mid := (start+end) >> 1
		if nums[mid] == target {
			// if target exists from the start or we reach target start point
			if mid == 0 || nums[mid-1] < target{
				return mid
			}
			if nums[mid-1] == target {
				end = mid-1
			}
		} else if nums[mid] < target {
			start = mid+1
		}
	}
	return -1
}

// finds last occurence of target in non-decreasing array where target is least elemt 
func findLastOccurence(nums []int, start, end, target int) int {
	for start <= end {
		mid := (start+end) >> 1
		if nums[mid] == target {
			// if target exists from the start or we reach target start point
			if mid == len(nums)-1 || nums[mid+1] > target{
				return mid
			}
			if nums[mid+1] == target {
				start = mid+1
			}
		} else if nums[mid] > target {
			end = mid-1
		}
	}
	return -1
}