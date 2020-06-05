package main

// T: O(n ^ 2) // n is the size of matrix
// M: O(n)
// -- start --

func isValidSudoku(board [][]byte) bool {

	// rows
	for i := 0; i < 9; i++ {
		m := [10]byte{}

		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			if m[board[i][j]-'0'] == 1 {
				return false
			}

			m[board[i][j]-'0']++
		}
	}

	// cols
	for i := 0; i < 9; i++ {
		m := [10]byte{}

		for j := 0; j < 9; j++ {
			if board[j][i] == '.' {
				continue
			}

			if m[board[j][i]-'0'] == 1 {
				return false
			}

			m[board[j][i]-'0']++
		}
	}

	// sub box
	for x := 0; x < 9; x += 3 {
		for y := 0; y < 9; y += 3 {
			m := [10]byte{}

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if board[x+i][y+j] == '.' {
						continue
					}

					if m[board[x+i][y+j]-'0'] == 1 {
						return false
					}

					m[board[x+i][y+j]-'0']++
				}
			}
		}
	}

	return true
}

// -- end --
