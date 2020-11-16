package main

func main() {

}

func gameOfLife(board [][]int) {

	if board == nil || len(board) == 0 || len(board[0]) == 0 {
		return
	}

	row, col := len(board), len(board[0])
	res := make([][]int, row)
	for i := range res {
		res[i] = make([]int, col)
	}

	x := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	y := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			count := 0
			for k := 0; k < 8; k++ {
				x := i + x[k]
				y := j + y[k]
				if x >= 0 && y >= 0 && x < row && y < col {
					count = count + (board[x][y])
				}
			}

			if board[i][j] == 1 {
				if count < 2 {
					res[i][j] = 0
				} else if count == 2 || count == 3 {
					res[i][j] = 1
				} else if count > 3 {
					res[i][j] = 0
				}
			} else {

				if count == 3 {
					res[i][j] = 1
				}
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			board[i][j] = res[i][j]
		}
	}
}
