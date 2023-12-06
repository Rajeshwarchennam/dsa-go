package MergeSortedArrays

// link - https://leetcode.com/problems/merge-sorted-array/description/

// nums1 and nums2 are sorted with sizes m, n resp
// nums1 has trailing zeroes of size n to accomadate the merged array 
func Merge(nums1 []int, nums2 []int, m, n int) {
	// i, j nums1, nums2 pointers  to exisiting elements
	// k nums1 pointer that fills elements from end
	i, j, k := m-1, n-1, m+n-1
	// only checking for nums2 pointer because 
	// if j becomes zero and i isn't zero yet 
	// we can ignore because we need to replace those elements 
	// in nums1 itself
	for j >= 0 {
		if i >= 0 {
			// if nums1 elements exist, we compare with nums2 element
			// and put that element which is highest at k pointer 
			// and reduce the pointer from whichever array we picked
			if nums1[i] > nums2[j] {
				nums1[k] = nums1[i]
				i--
			} else {
				nums1[k] = nums2[j]
				j--
			}
		} else {
			// if nums1 elements are done, 
			// we just put nums2 elements
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}