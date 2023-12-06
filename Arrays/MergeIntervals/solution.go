package MergeIntervals

// link - https://leetcode.com/problems/merge-intervals/description/

import "sort"

func Merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n == 0 {
		return nil
	}
	//sort the intervals on start times in asc order
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//this contains mergedIntervals
	mergedIntervals := make([][]int, 0, n)
	//helps in finding merged interval at current traversing
	currIntervalStart, currIntervalEnd := intervals[0][0], intervals[0][1]

	for i := 1; i<n; i++ {
		// if new interval is not overlapping with currentInterval, we can append current
		// interval to mergedIntervals and start new interval instance
		if intervals[i][0] > currIntervalEnd {
			mergedIntervals = append(mergedIntervals, []int{currIntervalStart, currIntervalEnd})
			currIntervalStart, currIntervalEnd = intervals[i][0], intervals[i][1]
		} else {
			// if overlapping, just update the currIntervalEnd to whatever max ending 
			currIntervalEnd = max(currIntervalEnd, intervals[i][1])
		}
	}

	// at the end of interation there will be one mergedInterval that is not appended to results
	// hence, appending that
	mergedIntervals = append(mergedIntervals, []int{currIntervalStart, currIntervalEnd})
	return mergedIntervals
}

func max(a, b int) int {
	if a>b {return a}
	return b
}