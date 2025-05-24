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
		return uint32ToRLColors(C_GRID)
	case 1:
		return uint32ToRLColors(C_T1)
	case 2:
		return uint32ToRLColors(C_T2)
	case 3:
		return uint32ToRLColors(C_T3)
	case 4:
		return uint32ToRLColors(C_T4)
	case 5:
		return uint32ToRLColors(C_T5)
	case 6:
		return uint32ToRLColors(C_T6)
	case 7:
		return uint32ToRLColors(C_T7)
	case 8:
		return uint32ToRLColors(C_T8)
	default:
		return uint32ToRLColors(C_GRID)
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

func rotateRight(matrix *[4][4]int, g *Game, xPosition int, yPosition int) (bool, int, int) {
	var rotated [4][4]int

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			rotated[j][3-i] = matrix[i][j]
		}
	}
	fmt.Println("Rotated:")
	fmt.Println(rotated)
	// Push down if bottom rows are empty (cut and drop)
	emptyRows := 0
	for i := 3; i >= 0; i-- {
		rowEmpty := true
		for j := 0; j < 4; j++ {
			if rotated[i][j] != 0 {
				rowEmpty = false
				break
			}
		}
		if rowEmpty {
			emptyRows++
		} else {
			break
		}
	}
	if emptyRows > 0 {
		for i := 3; i >= emptyRows; i-- {
			for j := 0; j < 4; j++ {
				rotated[i][j] = rotated[i-emptyRows][j]
			}
		}
		for i := 0; i < emptyRows; i++ {
			for j := 0; j < 4; j++ {
				rotated[i][j] = 0
			}
		}
	}

	type offset struct{ dx, dy int }
	kicks := []offset{{0, 0}, {1, 0}, {-1, 0}}

	for _, kick := range kicks {
		collision := false
		for i := 0; i < 4 && !collision; i++ {
			for j := 0; j < 4; j++ {
				if rotated[i][j] == 0 {
					continue
				}
				y := i + kick.dy + yPosition
				x := j + kick.dx + xPosition
				if y < 0 || y >= len(g.copyTable) || x < 0 || x >= len(g.copyTable[0]) {
					collision = true
					continue
				}
				if g.copyTable[y][x] != 0 {
					collision = true
					break
				}
			}
		}
		if !collision {
			// Apply rotated matrix
			for i := 0; i < 4; i++ {
				for j := 0; j < 4; j++ {
					matrix[i][j] = rotated[i][j]
				}
			}
			return true, kick.dx, kick.dy
		}
	}

	return false, 0, 0
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

func deleteCompleteLines(g *Game) {
	numberOfRows := 0
	for i := len(g.copyTable) - 1; i >= 0; i-- {
		complete := true
		for _, v := range g.copyTable[i] {
			if v == 0 {
				complete = false
				break
			}
		}
		if complete {
			for j := i; j > 0; j-- {
				g.copyTable[j] = g.copyTable[j-1]
			}
			g.copyTable[0] = [10]int{}
			i++
			numberOfRows++
		}
	}
	if numberOfRows > 0 {
		// 10 0 2 4
		// 10 0

		g.combo += numberOfRows
		score := SCORE_ROW * (g.combo * 2)
		g.score += score
	} else {
		g.combo = 0
	}
}
