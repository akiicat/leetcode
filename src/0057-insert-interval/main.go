package main

// T: O(n) // n elements
// M: O(1)
// -- start --

func insert(intervals [][]int, newInterval []int) [][]int {

	res := make([][]int, 0)

	i := 0

	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}

	// if overlap update new Interval
	for i < len(intervals) && intervals[i][0] <= newInterval[1] && newInterval[0] <= intervals[i][1] {
		newInterval[0] = Min(intervals[i][0], newInterval[0])
		newInterval[1] = Max(intervals[i][1], newInterval[1])
		i++
	}

	res = append(res, newInterval)

	for i < len(intervals) {
		res = append(res, intervals[i])
		i++
	}

	return res
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// -- end --
