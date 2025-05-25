package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREENWIDTH           = 630
	SCREENHEIGHT          = 840
	G_TALL                = 24
	G_WIDE                = 10
	G_START_X             = SCREENWIDTH/2 - SCREENWIDTH/4
	G_START_Y             = 40
	G_CELL_WIDTH          = 30
	G_CELL_HEIGHT         = 30
	FALL_SPEED            = 8
	SCORE_ROW             = 10
	TEXT_SCORE_POSITION_X = 15
	TEXT_SCORE_POSITION_Y = 20
	COYOTE_TIME           = 2
)

// EACH tetromino will be its color, 0 will be empty
type Tetromino struct {
	piece [4][4]int
	x     int
	y     int
}

type Game struct {
	isRunning   bool
	copyTable   [G_TALL][G_WIDE]int
	grid        [G_TALL][G_WIDE]int
	count       int
	coyoteTime  int
	onCollision bool
	tetromino   Tetromino
	score       int
	combo       int
}

func main() {
	fmt.Println("init")
	rl.InitWindow(SCREENWIDTH, SCREENHEIGHT, "Tetris")

	g := Game{coyoteTime: 0, isRunning: true, count: 0, tetromino: Tetromino{piece: tetrominos[0], x: 0, y: 0}}
	rl.SetTargetFPS(60)

	col := C_BACKGROUND
	fmt.Println(col)
	fmt.Println(uint32ToRLColors(col))

	g.gameNew()
	for !rl.WindowShouldClose() {
		if g.isRunning == false {
			break
		}

		getInput(&g)

		if g.count > FALL_SPEED {
			canFall(&g, g.tetromino.x, g.tetromino.y)

			if g.onCollision {
				g.coyoteTime++
				if g.coyoteTime > COYOTE_TIME {
					g.gameNew()
				}
			} else {
				moveTetromino(&g, g.tetromino.x, g.tetromino.y+1)
			}
			g.count = 0
		}
		g.count++
		rl.BeginDrawing()
		rl.ClearBackground(uint32ToRLColors(C_BACKGROUND))
		drawGrid(&g)
		g.drawHUD()
		rl.EndDrawing()
	}

	defer rl.CloseWindow()
}

func (g *Game) drawHUD() {
	rl.DrawRectangle(TEXT_SCORE_POSITION_X-10, G_START_Y, 140, 60, uint32ToRLColors(C_BACKGROUND))
	rl.DrawRectangleLines(TEXT_SCORE_POSITION_X-10, G_START_Y, 140, 60, uint32ToRLColors(C_BORDER))
	rl.DrawRectangleLines(G_START_X, G_START_Y, (G_CELL_WIDTH+1)*G_WIDE, (G_CELL_HEIGHT+1)*G_TALL, uint32ToRLColors(C_BORDER))
	scoreStr := fmt.Sprintf("SCORE: %d\nCOMBO: %d\n", g.score, g.combo)
	rl.DrawText(scoreStr, TEXT_SCORE_POSITION_X, G_START_Y+10, 20, uint32ToRLColors(C_TEXT))

}

func (g *Game) gameNew() {
	g.onCollision = false
	g.tetromino.piece = tetrominos[randomTetromino()]
	g.tetromino.y = 0
	g.tetromino.x = 4
	g.copyTable = g.grid
	checkForDeath(g)
	deleteCompleteLines(g)

}

func checkForDeath(g *Game) {
	for i := 0; i < 2; i++ {
		for j := range g.copyTable[i] {
			if g.copyTable[i][j] != 0 {
				g.isRunning = false
			}
		}

	}
}

func canFall(g *Game, x int, y int) {
	tetromino := g.tetromino.piece
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if tetromino[i][j] == 0 {
				continue
			}
			gridX := x + j
			gridY := y + i

			// Check if we're at the bottom or there's a block below
			if gridY+1 >= G_TALL || g.copyTable[gridY+1][gridX] != 0 {
				g.onCollision = true
				return
			}
		}
	}
	g.onCollision = false
}

func getInput(g *Game) {
	if rl.IsKeyPressed(rl.KeyQ) {
		printGrid(g)
		g.isRunning = false
	}

	if rl.IsKeyPressed(rl.KeyK) || rl.IsKeyPressed(rl.KeyRight) {
		moveTetromino(g, g.tetromino.x+1, g.tetromino.y)
	}

	if rl.IsKeyPressed(rl.KeyJ) || rl.IsKeyPressed(rl.KeyLeft) {
		moveTetromino(g, g.tetromino.x-1, g.tetromino.y)
	}

	if rl.IsKeyPressed(rl.KeySpace) {
		success, dx, dy := rotateRight(&g.tetromino.piece, g, g.tetromino.x, g.tetromino.y)
		if success {
			g.tetromino.x += dx
			g.tetromino.y += dy
			g.coyoteTime = 0
		}
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
// func canMove(piece [4][4]int, space [4][4]int) int {
// 	for i, row := range piece {
// 		for j := range row {
// 			if piece[i][j] != 0 && space[i][j] != 0 {
// 				return 1
// 			}
// 		}
// 	}
// 	return 0
// }
//
// func getSpaceFromPosition(g *Game, x int, y int) ([4][4]int, error) {
// 	var m [4][4]int
//
// 	for i := 0; i < 4; i++ {
// 		for j := 0; j < 4; j++ {
// 			gridY := y + i
// 			gridX := x + j
//
// 			if gridY < 0 || gridY >= len(g.grid) || gridX < 0 || gridX >= len(g.grid[0]) {
// 				m[i][j] = 1
// 			} else {
// 				m[i][j] = g.grid[gridY][gridX]
// 			}
// 		}
// 	}
//
// 	return m, nil
// }

func moveTetromino(g *Game, newX int, newY int) {
	g.onCollision = false
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
				g.onCollision = true
				fmt.Println("Vertical Collision")
				return
			}

			newCopy[y][x] = tetro.piece[i][j]
		}
	}
	g.tetromino.x = newX
	g.tetromino.y = newY
	g.grid = newCopy
	g.onCollision = false
}

func canPlace(rotated [4][4]int, grid [4][4]int, rowOffset, colOffset int) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			newRow := i + rowOffset
			newCol := j + colOffset
			if rotated[i][j] != 0 {
				if newRow < 0 || newRow >= 4 || newCol < 0 || newCol >= 4 {

					return false
				}
				fmt.Println("STATE OF GRID IN canPlace")
				fmt.Println(grid)
				if grid[newRow][newCol] != 0 {
					return false
				}
			}
		}
	}
	return true
}

func applyOffset(matrix *[4][4]int, rowOffset, colOffset int) {
	var moved [4][4]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if matrix[i][j] != 0 {
				moved[i+rowOffset][j+colOffset] = matrix[i][j]
			}
		}
	}
	*matrix = moved
}
