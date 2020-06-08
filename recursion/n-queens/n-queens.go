package main

//
//51. N皇后
//n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
//给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
//每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
//示例:
//输入: 4
//输出: [
//[".Q..",  // 解法 1
//"...Q",
//"Q...",
//"..Q."],
//
//["..Q.",  // 解法 2
//"Q...",
//"...Q",
//".Q.."]
//]
//解释: 4 皇后问题存在两个不同的解法。
//提示：
//皇后，是国际象棋中的棋子，意味着国王的妻子。皇后只做一件事，那就是“吃子”。
//当她遇见可以吃的棋子时，就迅速冲上去吃掉棋子。当然，她横、竖、斜都可走一到七步，可进可退。
//

var res [][]string

func solveNQueens(n int) [][]string {
	res = [][]string{}
	chessBoard := make([][]bool, n)
	for i := 0; i < n; i++ {
		chessBoard[i] = make([]bool, n)
	}
	trackBack(chessBoard, [][]byte{})
	return res
}

func trackBack(chessBoard [][]bool, track [][]byte) {
	if len(track) == len(chessBoard) {
		t := make([]string, len(track))
		for k, bs := range track {
			t[k] = string(bs)
		}
		res = append(res, t)
	}

	for j := 0; j < len(chessBoard); j++ {
		if !valid(chessBoard, len(track), j) { // 减支
			continue
		}
		bs := getLine(len(chessBoard))
		bs[j] = 'Q'
		chessBoard[len(track)][j] = true //棋盘放子
		track = append(track, bs)        // 加入选择路径
		trackBack(chessBoard, track)     // 递归
		track = track[:len(track)-1]     // 回溯
		chessBoard[len(track)][j] = false
	}
}

func valid(chessBoard [][]bool, row, cow int) bool {
	var i, j int
	for i = 0; i < row; i++ {
		if chessBoard[i][cow] == true {
			return false
		}
	}

	for j = 0; j < len(chessBoard); j++ {
		if chessBoard[row][j] == true {
			return false
		}
	}

	i, j = row, cow
	for i >= 0 && j >= 0 {
		if chessBoard[i][j] == true {
			return false
		}
		i--
		j--
	}

	i, j = row, cow
	for i >= 0 && j < len(chessBoard) {
		if chessBoard[i][j] == true {
			return false
		}
		j++
		i--
	}
	return true
}

func getLine(n int) []byte {
	bs := make([]byte, n)
	for i := 0; i < n; i++ {
		bs[i] = '.'
	}
	return bs
}
