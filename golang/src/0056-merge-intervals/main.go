package main

import "sort"

// T: O(n*log(n)) // n elements
// M: O(1)
// -- start --

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	rtn := [][]int{intervals[0]}
	rtn_size := 1

	for _, val := range intervals {
		if rtn[rtn_size-1][1] >= val[0] {
			rtn[rtn_size-1][1] = Max(rtn[rtn_size-1][1], val[1])
		} else {
			rtn = append(rtn, val)
			rtn_size++
		}
	}

	return rtn
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// -- end --
