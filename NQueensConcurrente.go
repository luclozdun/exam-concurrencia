package main

import(
	"fmt"
	"sync"
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

func nqueens(board []int, row int, doneChan chan struct{}){
	n := len(board)
  	if row == n{
		fmt.Println(board)
		doneChan <- struct{}{}
	} else{
		for col := 0; col < n; col++{
			if valid(board, row, col){
				board[row] = col
				nqueens(board, row+1, doneChan)
			}
		}
	}
}

func main(){

	var mu sync.Mutex

	doneChan := make(chan struct{})

	n := 5

	
	for i := 0; i < n; i++{
		board := []int{}
		for i:=0; i < n; i++{
			board = append(board, -1)
		}
		board[0] = i		
		mu.Lock()
		go nqueens(board, 1, doneChan)
		mu.Unlock()	
	}
	

	finish := 0
	for finish < n{
		<- doneChan
		finish++
	}
	fmt.Println("Terminado")
}