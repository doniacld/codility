package main

// [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]
/*
func main() {
	board := [][]string{{"A", "B"}, {"C", "A"}}
	word := "AC"

	exist(board, word)
}
*/
func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] && dfs(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i, j, index int, word string) bool {
	// the end of the word !
	if index == len(word) {
		return true
	}

	// extreme cases
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) ||  board[i][j] != word[index]  {
		return false
	}

/*
	// it's not a match, let's continue!
	if board[i][j] != word[index] {
		return false
	}
*/
	tmp := board[i][j]
	board[i][j] = ' '
	// if computed separately, it takes too much time!!
	found := dfs(board, i+1, j, index+1, word) ||  dfs(board, i-1, j, index+1, word) || dfs(board, i, j+1, index+1, word) || dfs(board, i, j-1, index+1, word)

	// reset the tmp value into the board place
	board[i][j] = tmp

	return found
}

/*
func exist(board [][]byte, word string) bool {
    for i:=0; i < len(board); i++ {
        for j:=0; j < len(board[0]); j++ {
            if board[i][j] == word[0] && dfs(board, i, j, 0, word) {
                return true
            }
        }
    }
    return false
}

func dfs(board [][]byte, i int, j int, idx int, word string) bool {
    if idx == len(word) { return true }
    if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) ||
            board[i][j] != word[idx] { return false }

    temp:= board[i][j]
    board[i][j] = ' '
    found := dfs(board, i+1, j, idx+1, word) || dfs(board, i-1, j, idx+1, word) ||
        dfs(board, i, j+1, idx+1, word) || dfs(board, i, j-1, idx+1, word)
    board[i][j] = temp
    return found
}

*/