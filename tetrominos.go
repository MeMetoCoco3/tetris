package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

// var tetrominos [8][4][4]int = [8][4][4]int{
// 	{
// 		{0, 0, 0, 0},
// 		{0, 0, 0, 0},
// 		{0, 0, 0, 0},
// 		{0, 0, 0, 0},
// 	},
// 	{
// 		{9, 1, 9, 9},
// 		{9, 1, 9, 9},
// 		{9, 1, 9, 9},
// 		{9, 1, 9, 9},
// 	},
// 	{
// 		{9, 2, 9, 9},
// 		{9, 2, 9, 9},
// 		{9, 2, 2, 9},
// 		{9, 9, 9, 9},
// 	},
// 	{
// 		{9, 9, 3, 9},
// 		{9, 9, 3, 9},
// 		{9, 3, 3, 9},
// 		{9, 9, 9, 9},
// 	},
// 	{
// 		{9, 4, 9, 9},
// 		{9, 4, 4, 9},
// 		{9, 4, 9, 9},
// 		{9, 9, 9, 9},
// 	},
// 	{
// 		{9, 9, 9, 9},
// 		{9, 5, 5, 9},
// 		{9, 5, 5, 9},
// 		{9, 9, 9, 9},
// 	},
// 	{
// 		{9, 9, 9, 9},
// 		{9, 6, 6, 9},
// 		{6, 6, 9, 9},
// 		{9, 9, 9, 9},
// 	},
// 	{
// 		{9, 9, 9, 9},
// 		{9, 7, 7, 9},
// 		{9, 9, 7, 7},
// 		{9, 9, 9, 9},
// 	},
// }

var tetrominos [8][4][4]int = [8][4][4]int{
	{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	},
	{
		{0, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 1, 0, 0},
	},
	{
		{0, 0, 0, 0},
		{0, 2, 0, 0},
		{0, 2, 0, 0},
		{0, 2, 2, 0},
	},
	{
		{0, 0, 0, 0},
		{0, 0, 3, 0},
		{0, 0, 3, 0},
		{0, 3, 3, 0},
	},
	{
		{0, 0, 0, 0},
		{0, 4, 0, 0},
		{0, 4, 4, 0},
		{0, 4, 0, 0},
	},
	{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 5, 5, 0},
		{0, 5, 5, 0},
	},
	{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 6, 6, 0},
		{6, 6, 0, 0},
	},
	{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 7, 7, 0},
		{0, 0, 7, 7},
	},
}

func printGrid(g *Game) {
	for i, _ := range g.grid {
		fmt.Printf("| ")
		for _, vj := range g.grid[i] {
			fmt.Printf("%d ", vj)
		}
		fmt.Printf("|\n")

	}

}

func getCellColor(i int) rl.Color {
	switch i {
	case 0:
		return rl.Black
	case 1:
		return rl.Red
	case 2:
		return rl.Blue
	case 3:
		return rl.Yellow
	case 4:
		return rl.Green
	case 5:
		return rl.Purple
	case 6:
		return rl.Orange
	case 7:
		return rl.Pink
	case 8:
		return rl.Brown
	default:
		return rl.Black
	}
}

func randomTetromino() int {
	return (rand.Intn(7) + 1)

}

func drawTetrominoOnGrid(x int, y int, t [4][4]int, g *Game) {
	for i, it := range t {
		for j, jt := range it {
			if jt == 0 {
				continue
			}
			// if g.grid[y+i][x+j] == 0 {
			if x+j == -1 {
				x++
			}
			if x+j == 10 {
				x--
			}
			g.grid[y+i][x+j] = jt
		}
	}
}

// }

func rotateRight(matrix *[4][4]int, positionOnGrid [4][4]int) {
	var rotated [4][4]int
	// Copy = rotated matrix
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			rotated[j][3-i] = matrix[i][j]
		}
	}

	// If we have a bottom line with 0s we cut it and push down the tetromino
	empty_rows := 0
	for i := 3; i >= 0; i-- {
		inline_zeroes := 0
		for j := 0; j < 4; j++ {
			if rotated[i][j] == 0 {
				inline_zeroes++
			}
		}
		if inline_zeroes == 4 {
			empty_rows++
		} else {
			break
		}
	}

	if empty_rows > 0 {
		for i := 3; i >= empty_rows; i-- {
			for j := 0; j < 4; j++ {
				rotated[i][j] = rotated[i-empty_rows][j]
			}
		}
		for i := 0; i < empty_rows; i++ {
			for j := 0; j < 4; j++ {
				rotated[i][j] = 0
			}
		}
	}

	isOk := 1
	tries := []int{0}
	for isOk > 0 {
		isOk = 0
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				// maybe out of bounds if rotate on right side
				if positionOnGrid[i][j] != 0 && rotated[i][j] != 0 {
					isOk = 1
				}
			}
		}

		if isOk != 0 {
			rightIndex := shiftTetrominoRight(&rotated)
			tries = append(tries, rightIndex)
			isOk = 0
			if tries[len(tries)-1] != 0 {
				shiftTetrominoLeft(&rotated)
				shiftTetrominoLeft(&rotated)
				isOk = -1
			}
		}
	}

	if isOk == 0 {
		// matrix = copy
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				matrix[i][j] = rotated[i][j]
			}
		}
	}
}

func shiftTetrominoRight(tetromino *[4][4]int) int {
	rightestValue := 0
	for i := range tetromino {
		for vIndex, v := range tetromino[i] {
			if v != 0 && vIndex >= rightestValue {
				rightestValue = vIndex
			}
		}
	}
	if rightestValue > 0 {
		for i := 0; i < 4; i++ {
			for j := 3; j > 0; j-- {
				tetromino[i][j] = tetromino[i][j-1]
			}
			tetromino[i][0] = 0
		}
	}
	return rightestValue
}

func shiftTetrominoLeft(tetromino *[4][4]int) int {
	leftestValue := 3
	for i := range tetromino {
		for vIndex, v := range tetromino[i] {
			if v != 0 && vIndex <= leftestValue {
				leftestValue = vIndex
			}
		}
	}

	if leftestValue < 3 {
		for i := 0; i < 4; i++ {
			for j := 0; j < 3; j++ {
				tetromino[i][j] = tetromino[i][j+1]
			}
			tetromino[i][3] = 0
		}
	}
	return leftestValue
}

func copyTetromino(tetromino *[4][4]int) *[4][4]int {
	var copy [4][4]int
	for i := range tetromino {
		for j := range tetromino[i] {
			copy[i][j] = tetromino[i][j]
		}
	}
	return &copy
}
