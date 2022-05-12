package main

import(
	"fmt"
)

func valid(board []int, row, col int)(bool){
	for row_i := 0; row_i < row; row_i++{
		col_i := board[row_i]
	    delta := row - row_i
    	if col == col_i{
			return false
		} else if col == (col_i + delta){
			return false
		} else if col == (col_i - delta){
			return false
		}
	}    
  	return true
}

func nqueens(board []int, row int){
	n := len(board)
  	if row == n{
		fmt.Println(board)
	} else{
		for col := 0; col < n; col++{
			if valid(board, row, col){
				board[row] = col
				nqueens(board, row+1)
			}
		}
	}
}

func main(){
	n := 4
	
	board := []int{}
	for i:=0; i < n; i++{
		board = append(board, -1)
	}

	nqueens(board, 0)
}