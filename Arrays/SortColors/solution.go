package SortColors

// link - https://leetcode.com/problems/sort-colors/

// TC - O(N)
// SC - O(1)


func SortColors(nums []int) {

	n := len(nums)
	if n == 0 {
		return
	}

	// Dutch National flag algo

	// low, mid, high properties as follows
	// 0 to low-1 - 0s
	// low to mid-1 - 1s
	// mid to high - unsorted
	// high+1 to n-1 - 2s

	low, mid, high := 0, 0, n-1

	for mid <= high {
		switch nums[mid] {
		case 0:
			// if nums[mid] == 0, swap low, mid elements
			// increment low and mid
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			// increment mid if nums[mid]==1, as that satisfies above property
			mid++
		case 2:
			// swap mid, high and decrease high so that we statisfy above property
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
}