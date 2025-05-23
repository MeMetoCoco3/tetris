package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREENWIDTH   = 600
	SCREENHEIGHT  = 800
	G_TALL        = 20
	G_WIDE        = 10
	G_START_X     = SCREENWIDTH/2 - SCREENWIDTH/4
	G_START_Y     = 40
	G_CELL_WIDTH  = 30
	G_CELL_HEIGHT = 30
)

// EACH tetromino will be its color, 0 will be empty
type Tetromino struct {
	piece [4][4]int
	x     int
	y     int
}

// 0 dead 1 alive

type Game struct {
	isRunning   bool
	copyTable   [G_TALL][G_WIDE]int
	grid        [G_TALL][G_WIDE]int
	count       int
	onCollision int
	tetromino   Tetromino
}

func main() {
	fmt.Println("init")
	rl.InitWindow(SCREENWIDTH, SCREENHEIGHT, "Tetris")

	g := Game{isRunning: true, count: 0, tetromino: Tetromino{piece: tetrominos[0], x: 0, y: 0}}
	rl.SetTargetFPS(60)

	g.gameNew()
	for !rl.WindowShouldClose() {
		getInput(&g)
		if g.isRunning == false {
			break
		}

		if g.onCollision > 0 {
			g.onCollision++
			if g.onCollision > 5 {
				g.gameNew()
			}
		}

		if g.count > 10 {

			if g.tetromino.y == G_TALL-4 {
				g.gameNew()
			} else {
				moveTetromino(&g, g.tetromino.x, g.tetromino.y+1)
			}
			g.count = 0
		}
		g.count++

		rl.BeginDrawing()
		rl.ClearBackground(rl.Red)

		drawGrid(&g)

		rl.EndDrawing()
	}

	defer rl.CloseWindow()
}

func (g *Game) gameNew() {
	g.onCollision = 0
	g.tetromino.piece = tetrominos[randomTetromino()]
	g.tetromino.y = 0
	g.tetromino.x = 4
	g.copyTable = g.grid

}

func getInput(g *Game) {
	if rl.IsKeyPressed(rl.KeyQ) {
		printGrid(g)
		g.isRunning = false
	}

	if rl.IsKeyPressed(rl.KeyK) {
		moveTetromino(g, g.tetromino.x+1, g.tetromino.y)
	}

	if rl.IsKeyPressed(rl.KeyJ) {
		moveTetromino(g, g.tetromino.x-1, g.tetromino.y)
	}

	if rl.IsKeyPressed(rl.KeyBackspace) {
		rotateRight(&g.tetromino.piece)
	}

}

func cleanGrid(g *Game) {
	g.grid = g.copyTable
	//	for i := 0; i < G_TALL; i++ {
	//		for j := 0; j < G_WIDE; j++ {
	//			g.grid[i][j] = 0
	//		}
	//	}
}

func drawGrid(g *Game) {
	cleanGrid(g)
	piece := g.tetromino.piece
	x := g.tetromino.x
	y := g.tetromino.y
	// Todo: check if we can draw tetromino

	drawTetrominoOnGrid(x, y, piece, g)
	// for i := 0; i < len(piece); i++ {
	// 	for j := 0; j < len(piece[i]); j++ {
	// 		g.grid[i+g.tetromino.y][j+g.tetromino.x] = piece[i][j]
	//
	// 	}
	// }
	for i := 0; i < G_TALL; i++ {
		for j := 0; j < G_WIDE; j++ {
			color := getCellColor(g.grid[i][j])
			x := int32(G_START_X + j*(G_CELL_WIDTH+1))
			y := int32(G_START_Y + i*(G_CELL_HEIGHT+1))
			rl.DrawRectangle(x, y, G_CELL_WIDTH, G_CELL_HEIGHT, color)
		}
	}
}

// HELPERS, 1 IF IT HAS SOMETHING ON, 0 OTHERWISe
func canMove(piece [4][4]int, space [4][4]int) int {
	for i, row := range piece {
		for j := range row {
			if piece[i][j] != 0 && space[i][j] != 0 {
				return 1
			}
		}
	}
	return 0
}

func getTetrominoWidth(tetromino [4][4]int) int {
	minCol := 4
	maxCol := -1

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if tetromino[i][j] != 0 {
				if j < minCol {
					minCol = j
				}
				if j > maxCol {
					maxCol = j
				}
			}
		}
	}

	return maxCol - minCol + 1
}

func getSpaceFromPosition(g *Game, x int, y int, pieceWidth int) ([4][4]int, error) {
	m := [4][4]int{}

	if x < 0 || x+pieceWidth > len(g.grid[0]) {
		return m, fmt.Errorf("Cant Move To New Space")
	}

	y_index := 0
	for i := y; i < y+pieceWidth; i++ {
		x_index := 0
		for j := x; j < x+pieceWidth; j++ {
			m[y_index][x_index] = g.grid[i][j]
			x_index++
		}
		y_index++
	}

	return m, nil

}

func moveTetromino(g *Game, newX int, newY int) {
	g.onCollision = 0
	tetro := g.tetromino

	var newCopy [G_TALL][G_WIDE]int
	for i := range g.copyTable {
		for j := range g.copyTable[i] {
			newCopy[i][j] = g.copyTable[i][j]
		}
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if tetro.piece[i][j] == 0 {
				continue
			}
			x := j + newX
			y := i + newY

			if x < 0 || x >= G_WIDE {
				fmt.Println(" ERROR ON X")
				return
			}

			if y < 0 || y > G_TALL {
				fmt.Println(" ERROR ON Y")
				return
			}

			if newCopy[y][x] != 0 {
				if newY == g.tetromino.y {
					fmt.Println("Horizontal Collision")
					return
				}
				g.onCollision = 1
				fmt.Println("Vertical Collision")
				return
			}

			newCopy[y][x] = tetro.piece[i][j]
		}
	}
	g.tetromino.x = newX
	g.tetromino.y = newY
	g.grid = newCopy

}
