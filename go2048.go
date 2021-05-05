package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/nsf/termbox-go"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

// 定义棋盘尺寸
var (
	board [4][4]int
	size  = len(board)
	score = 0
	best = 0
	moves = 0
)

// 用于重置棋盘
func newBoard() {
	board = [4][4]int{}
}

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
	// 清屏
	cmd := exec.Command("cmd.exe", "/c", "cls")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()

	color.Set(color.FgYellow)
	fmt.Printf("\n\n\n\n\n\n\n\n                                       join the numbers and get to the 2048!\n")
	fmt.Printf("                                              ESC:quit | ENTER:restart\n")
	fmt.Printf("                                              play with arrow keys\n")

	fmt.Printf("                                              ")
	color.Set(color.FgWhite)
	color.Set(color.BgMagenta)
	fmt.Printf("MOVES ")
	color.Set(color.FgWhite)
	color.Set(color.BgCyan)
	fmt.Printf(" SCORE ")
	color.Set(color.FgWhite)
	color.Set(color.BgGreen)
	fmt.Printf("   BEST \n")
	color.Unset()

	fmt.Printf("                                              ")
	color.Set(color.FgWhite)
	color.Set(color.BgMagenta)
	fmt.Printf("%5d ",moves)
	color.Set(color.FgHiWhite)
	color.Set(color.BgCyan)
	fmt.Printf(" %5d ",score)
	color.Set(color.FgWhite)
	color.Set(color.BgGreen)
	fmt.Printf("  %5d \n",best)
	color.Unset()

	fmt.Printf("                                              ")
	color.Set(color.BgWhite)
	color.Set(color.FgWhite)
	fmt.Println("-+---+----+----+----+")
	color.Unset()
	for _, v := range board {
		for i, v2 := range v {
			if i == 0 {
				fmt.Printf("                                              ")
				color.Set(color.BgWhite)
				color.Set(color.FgWhite)
				fmt.Printf("+")
				color.Unset()
			}
			switch v2 {
			case 0:
				color.Set(color.FgBlack)
				fmt.Printf("%4d", v2)
				color.Unset()
			case 2:
				fmt.Printf("%4d", v2)
			case 4:
				color.Set(color.FgHiBlue)
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
				color.Set(color.FgHiMagenta)
				fmt.Printf("%4d", v2)
				color.Unset()
			case 64:
				color.Set(color.FgCyan)
				fmt.Printf("%4d", v2)
				color.Unset()
			case 128:
				color.Set(color.FgHiYellow)
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
			color.Set(color.BgWhite)
			color.Set(color.FgWhite)
			fmt.Printf("+")
			color.Unset()
		}
		fmt.Println()
		fmt.Printf("                                              ")
		color.Set(color.BgWhite)
		color.Set(color.FgWhite)
		fmt.Println("-+---+----+----+----+")
		color.Unset()
	}
}

func main() {
	// 引入termbox包，用来识别方向键
	_ = termbox.Init()
	defer termbox.Close()
start:
	newBoard()
	rand.Seed(time.Now().UnixNano())
	//在board随机位置生成起始数字
	i := rand.Int() % size
	j := rand.Int() % size
	board[i][j] = 2
	for {
		showBoard()
		score = 0
		switch termbox.PollEvent().Key {
		case termbox.KeyArrowUp:
			moveUp()
		case termbox.KeyArrowDown:
			moveDown()
		case termbox.KeyArrowLeft:
			moveLeft()
		case termbox.KeyArrowRight:
			moveRight()
		case termbox.KeyEsc:
			os.Exit(0)
		case termbox.KeyCtrlC:
			os.Exit(0)
		case termbox.KeyEnter:
			moves = 0
			goto start
		default:

		}
		if checkWin() {
			showBoard()
			color.Set(color.FgHiGreen)
			fmt.Println("                                                $$$$YOU WIN!$$$$")
			color.Unset()
			break
		}
		if checkLose() {
			showBoard()
			color.Set(color.FgRed)
			fmt.Println("                                                ####GAME OVER####")
			color.Unset()
			break
		}
	}
continueOrNot:
	switch termbox.PollEvent().Key {
	case termbox.KeyEsc:
		os.Exit(0)
	case termbox.KeyEnter:
		moves = 0
		score = 0
		goto start
	default:
		goto continueOrNot
	}
}

func moveDown() {
	// changed标志位用来判断执行本次操作是否对棋盘数字造成改变，若没有改变，则不生成随机数字，即本次操作为无效操作
	changed := false
	for i := 0; i < size-1; i++ {
		for j := 0; j < size; j++ {
			// 同一列上下两行都为0 则不执行任何操作
			if board[i][j] == 0 && board[i+1][j] == 0 {
				continue
			}
			// 上下两行数字相同，则向下合并，并将标志位置为true
			if board[i][j] == board[i+1][j] {
				board[i+1][j] = 2 * board[i][j]
				board[i][j] = 0
				changed = true
				//这里解决特殊情况，比如遍历到第二行，执行上面的代码将本列第二行的数字和第三行合并，
				//此时若不做以下判断，第一行的数字将不往下合并
				k := i
				for k > 0 {
					if board[k-1][j] != 0 {
						board[k][j] = board[k-1][j]
						board[k-1][j] = 0
					}
					k--
				}
			}
			// 下一行为0，则本行数字往下移动
			if board[i+1][j] == 0 {
				board[i+1][j] = board[i][j]
				board[i][j] = 0
				changed = true
				k := i
				for k > 0 {
					if board[k-1][j] != 0 {
						board[k][j] = board[k-1][j]
						board[k-1][j] = 0
					}
					k--
				}
			}
		}
	}
	if !changed {
		return
	}
	moves++
	// 因为这里是执行向下合并的操作，按照2048的游戏规则应该在第一行随机空位生成一个随机数
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

// 原理同moveDown
func moveUp() {
	changed := false
	for i := size - 1; i > 0; i-- {
		for j := 0; j < size; j++ {
			if board[i][j] == 0 && board[i-1][j] == 0 {
				continue
			}
			if board[i][j] == board[i-1][j] {
				board[i-1][j] = 2 * board[i][j]
				board[i][j] = 0
				changed = true
				k := i
				for k < size-1 {
					if board[k+1][j] != 0 {
						board[k][j] = board[k+1][j]
						board[k+1][j] = 0
					}
					k++
				}
			}
			if board[i-1][j] == 0 {
				board[i-1][j] = board[i][j]
				board[i][j] = 0
				changed = true
				k := i
				for k < size-1 {
					if board[k+1][j] != 0 {
						board[k][j] = board[k+1][j]
						board[k+1][j] = 0
					}
					k++
				}
			}
		}
	}
	if !changed {
		return
	}
	moves++
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

//原理同moveDown
func moveRight() {
	changed := false
	for j := 0; j < size-1; j++ {
		for i := 0; i < size; i++ {
			if board[i][j] == 0 && board[i][j+1] == 0 {
				continue
			}
			if board[i][j] == board[i][j+1] {
				board[i][j+1] = 2 * board[i][j]
				board[i][j] = 0
				changed = true
				k := j
				for k > 0 {
					if board[i][k-1] != 0 {
						board[i][k] = board[i][k-1]
						board[i][k-1] = 0
					}
					k--
				}
			}
			if board[i][j+1] == 0 {
				board[i][j+1] = board[i][j]
				board[i][j] = 0
				changed = true
				k := j
				for k > 0 {
					if board[i][k-1] != 0 {
						board[i][k] = board[i][k-1]
						board[i][k-1] = 0
					}
					k--
				}
			}
		}
	}
	if !changed {
		return
	}
	moves++
	zeroRow := make([]int, 0)
	for i := 0; i < size; i++ {
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

//原理同moveDown
func moveLeft() {
	changed := false
	for j := size - 1; j > 0; j-- {
		for i := 0; i < size; i++ {
			if board[i][j] == 0 && board[i][j-1] == 0 {
				continue
			}
			if board[i][j] == board[i][j-1] {
				board[i][j-1] = 2 * board[i][j]
				board[i][j] = 0
				changed = true
				k := j
				for k < size-1 {
					if board[i][k+1] != 0 {
						board[i][k] = board[i][k+1]
						board[i][k+1] = 0
					}
					k++
				}
			}
			if board[i][j-1] == 0 {
				board[i][j-1] = board[i][j]
				board[i][j] = 0
				changed = true
				k := j
				for k < size-1 {
					if board[i][k+1] != 0 {
						board[i][k] = board[i][k+1]
						board[i][k+1] = 0
					}
					k++
				}
			}
		}
	}
	if !changed {
		return
	}
	moves++
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

//计分并判断是否获胜
func checkWin() bool {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			score = score + board[i][j]*(1+board[i][j]/32)
			if score>best {
				best = score
			}
			if board[i][j] == 2048 {
				return true
			}
		}
	}
	return false
}

//棋盘全部填满且不能合并则游戏结束
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
	return true
}
