package main

// T: O(n*m))
// M: O(1)
// -- start --

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	colSize := len(obstacleGrid)
	if colSize == 0 {
		return 0
	}
	rowSize := len(obstacleGrid[0])

	if obstacleGrid[0][0] == 1 {
		return 0
	}

	// set obstacles
	for col := 0; col < colSize; col++ {
		for row := 0; row < rowSize; row++ {
			if obstacleGrid[col][row] == 1 {
				obstacleGrid[col][row] = -1
			}
		}
	}

	// set start point
	obstacleGrid[0][0] = 1

	for col := 0; col < colSize; col++ {
		for row := 0; row < rowSize; row++ {
			if obstacleGrid[col][row] != -1 {
				if row-1 >= 0 && obstacleGrid[col][row-1] != -1 {
					obstacleGrid[col][row] += obstacleGrid[col][row-1]
				}

				if col-1 >= 0 && obstacleGrid[col-1][row] != -1 {
					obstacleGrid[col][row] += obstacleGrid[col-1][row]
				}
			}
		}
	}

	if obstacleGrid[colSize-1][rowSize-1] == -1 {
		return 0
	}

	return obstacleGrid[colSize-1][rowSize-1]
}

// -- end --
