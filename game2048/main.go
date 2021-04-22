package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"os"
	"strings"
	"time"
)

var board [5][5]int
var size = len(board)

//控制随机数的概率
func randNumber() int {
	n := rand.Intn(10)
	switch n {
	case 1, 2, 3, 4, 5:
		return 2
	case 6, 7, 8:
		return 4
	case 9:
		return 8
	}
	return 16
}

//输出board并配色
func showBoard() {
	fmt.Println("----+----+----+----+----+")
	for _, v := range board {
		for _, v2 := range v {
			switch v2 {
			case 0:
				fmt.Printf("%4d", v2)
			case 2:
				color.Set(color.FgHiMagenta)
				fmt.Printf("%4d", v2)
				color.Unset()
			case 4:
				color.Set(color.FgBlue)
				fmt.Printf("%4d", v2)
				color.Unset()
			case 8:
				color.Set(color.FgYellow)
				fmt.Printf("%4d", v2)
				color.Unset()
			case 16:
				color.Set(color.FgGreen)
				fmt.Printf("%4d", v2)
				color.Unset()
			case 32:
				color.Set(color.FgRed)
				fmt.Printf("%4d", v2)
				color.Unset()
			case 64:
				color.Set(color.FgCyan)
				fmt.Printf("%4d", v2)
				color.Unset()
			case 128:
				color.Set(color.FgRed, color.FgYellow)
				fmt.Printf("%4d", v2)
				color.Unset()
			case 256:
				color.Set(color.FgHiCyan)
				fmt.Printf("%4d", v2)
				color.Unset()
			case 512:
				color.Set(color.FgHiGreen)
				fmt.Printf("%4d", v2)
				color.Unset()
			case 1024:
				color.Set(color.FgHiBlue)
				fmt.Printf("%4d", v2)
				color.Unset()
			case 2048:
				color.Set(color.FgHiRed)
				fmt.Printf("%4d", v2)
				color.Unset()
			}
			fmt.Printf("|")
		}
		fmt.Println()
		fmt.Println("----+----+----+----+----+")
	}
}

func main() {
	var input string
	rand.Seed(time.Now().UnixNano())

	//在board随机位置生成起始数字
	i := rand.Int() % size
	j := rand.Int() % size
	board[i][j] = 2
	for {
		showBoard()
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		switch input {
		case "w":
			moveUp()
		case "s":
			moveDown()
		case "a":
			moveLeft()
		case "d":
			moveRight()
		case "q":
			os.Exit(0)
		}
		if checkWin() {
			break
		}
		if checkLose() {
			break
		}
	}
	fmt.Println("请按回车键退出")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func moveDown() {
	for i := 0; i < size-1; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == board[i+1][j] {
				board[i+1][j] = 2 * board[i][j]
				board[i][j] = 0
			}
			if board[i+1][j] == 0 {
				board[i+1][j] = board[i][j]
				board[i][j] = 0
			}
		}
	}
	zeroCol := make([]int, 0)
	for j := 0; j < size; j++ {
		if board[0][j] == 0 {
			zeroCol = append(zeroCol, j)
		}
	}
	if len(zeroCol) == 0 {
		return
	}
	randCol := zeroCol[rand.Int()%len(zeroCol)]
	board[0][randCol] = randNumber()
}

func moveUp() {
	for i := size - 1; i > 0; i-- {
		for j := 0; j < size; j++ {
			if board[i][j] == board[i-1][j] {
				board[i-1][j] = 2 * board[i][j]
				board[i][j] = 0
			}
			if board[i-1][j] == 0 {
				board[i-1][j] = board[i][j]
				board[i][j] = 0
			}
		}
	}
	zeroCol := make([]int, 0)
	for j := 0; j < size; j++ {
		if board[size-1][j] == 0 {
			zeroCol = append(zeroCol, j)
		}
	}
	if len(zeroCol) == 0 {
		return
	}
	randCol := zeroCol[rand.Int()%len(zeroCol)]
	board[size-1][randCol] = randNumber()
}

func moveRight() {
	for i := 0; i < size; i++ {
		for j := 0; j < size-1; j++ {
			if board[i][j] == board[i][j+1] {
				board[i][j+1] = 2 * board[i][j]
				board[i][j] = 0
			}
			if board[i][j+1] == 0 {
				board[i][j+1] = board[i][j]
				board[i][j] = 0
			}
		}
	}
	zeroRow := make([]int, 0)
	for i := 0; i < size-1; i++ {
		if board[i][0] == 0 {
			zeroRow = append(zeroRow, i)
		}
	}
	if len(zeroRow) == 0 {
		return
	}
	randCol := zeroRow[rand.Int()%len(zeroRow)]
	board[randCol][0] = randNumber()
}

func moveLeft() {
	for i := 0; i < size; i++ {
		for j := size - 1; j > 0; j-- {
			if board[i][j] == board[i][j-1] {
				board[i][j-1] = 2 * board[i][j]
				board[i][j] = 0
			}
			if board[i][j-1] == 0 {
				board[i][j-1] = board[i][j]
				board[i][j] = 0
			}
		}
	}
	zeroRow := make([]int, 0)
	for i := 0; i < size; i++ {
		if board[i][size-1] == 0 {
			zeroRow = append(zeroRow, i)
		}
	}
	if len(zeroRow) == 0 {
		return
	}
	randCol := zeroRow[rand.Int()%len(zeroRow)]
	board[randCol][size-1] = randNumber()
}

func checkWin() bool {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == 2048 {
				showBoard()
				fmt.Println("YOU WIN!")
				return true
			}
		}
	}
	return false
}

func checkLose() bool {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == 0 {
				return false
			}
		}
	}

	for i := 0; i < size-1; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == board[i+1][j] {
				return false
			}
		}
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size-1; j++ {
			if board[i][j] == board[i][j+1] {
				return false
			}
		}
	}
	fmt.Println("GAME OVER!")
	return true
}
