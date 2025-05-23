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
			fmt.Println(y+i, x+j, jt)
			g.grid[y+i][x+j] = jt
		}
	}
}

// }

func rotateRight(matrix *[4][4]int) {
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			matrix[i][j], matrix[i][4-1-j] = matrix[i][4-1-j], matrix[i][j]
		}
	}
	// If we have a bottom line with 0s we cut it and push down the tetromino
	empty_rows := 0
	for i := 3; i >= 0; i-- {
		inline_zeroes := 0
		for j := 0; j < 4; j++ {
			if matrix[i][j] == 0 {
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
				matrix[i][j] = matrix[i-empty_rows][j]
			}
		}
		// Clear top rows
		for i := 0; i < empty_rows; i++ {
			for j := 0; j < 4; j++ {
				matrix[i][j] = 0
			}
		}
	}

}
