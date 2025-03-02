package main

import (
	"image/color"
	"log"
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
	paddleHeight = 12
	paddleWidth  = 60
	ballSize     = 8
	brickWidth   = 60
	brickHeight  = 20
	brickRows    = 5
	brickCols    = 10
	brickGap     = 4
	paddleSpeed  = 5
	ballSpeed    = 3
	paddleY      = screenHeight - 40
	ballStartY   = screenHeight - 60
	brickOffset  = 40
)

type Game struct {
	paddle      paddle
	ball        ball
	bricks      []brick
	score       int
	gameOver    bool
	initialized bool
}

type paddle struct {
	x, y, width, height float64
}

type ball struct {
	x, y, dx, dy, size float64
}

type brick struct {
	x, y, width, height float64
	active              bool
	color               color.Color
	score               int
}

func (g *Game) init() {
	// Initialize paddle
	g.paddle = paddle{
		x:      screenWidth/2 - paddleWidth/2,
		y:      paddleY,
		width:  paddleWidth,
		height: paddleHeight,
	}

	// Initialize ball
	g.ball = ball{
		x:    screenWidth / 2,
		y:    ballStartY,
		dx:   ballSpeed,
		dy:   -ballSpeed,
		size: ballSize,
	}

	// Initialize bricks
	g.initBricks()
	g.initialized = true
}

func (g *Game) initBricks() {
	brickConfig := []struct {
		color color.Color
		score int
	}{
		{color.RGBA{0xff, 0x00, 0x00, 0xff}, 50}, // Red
		{color.RGBA{0xff, 0x7f, 0x00, 0xff}, 40}, // Orange
		{color.RGBA{0xff, 0xff, 0x00, 0xff}, 30}, // Yellow
		{color.RGBA{0x00, 0xff, 0x00, 0xff}, 20}, // Green
		{color.RGBA{0x00, 0x00, 0xff, 0xff}, 10}, // Blue
	}

	g.bricks = make([]brick, 0, brickRows*brickCols)
	for row := 0; row < brickRows; row++ {
		for col := 0; col < brickCols; col++ {
			g.bricks = append(g.bricks, brick{
				x:      float64(col*(brickWidth+brickGap) + brickGap),
				y:      float64(row*(brickHeight+brickGap) + brickGap + brickOffset),
				width:  brickWidth,
				height: brickHeight,
				active: true,
				color:  brickConfig[row].color,
				score:  brickConfig[row].score,
			})
		}
	}
}

func (g *Game) Update() error {
	if !g.initialized {
		g.init()
	}

	if g.gameOver {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.init()
			g.gameOver = false
			g.score = 0
		}
		return nil
	}

	// Update paddle position based on input
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.paddle.x -= paddleSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.paddle.x += paddleSpeed
	}

	// Keep paddle within screen bounds
	if g.paddle.x < 0 {
		g.paddle.x = 0
	}
	if g.paddle.x > screenWidth-g.paddle.width {
		g.paddle.x = screenWidth - g.paddle.width
	}

	// Update ball position
	g.ball.x += g.ball.dx
	g.ball.y += g.ball.dy

	// Ball collision with walls
	if g.ball.x <= 0 || g.ball.x >= screenWidth-g.ball.size {
		g.ball.dx = -g.ball.dx
	}
	if g.ball.y <= 0 {
		g.ball.dy = -g.ball.dy
	}

	// Ball collision with paddle
	if g.ball.y+g.ball.size >= g.paddle.y &&
		g.ball.x+g.ball.size >= g.paddle.x &&
		g.ball.x <= g.paddle.x+g.paddle.width &&
		g.ball.y <= g.paddle.y+g.paddle.height {
		g.ball.dy = -g.ball.dy
		// Add some randomness to the ball direction
		g.ball.dx += (rand.Float64()*2 - 1) * 0.5
	}

	// Ball collision with bricks
	for i := range g.bricks {
		if !g.bricks[i].active {
			continue
		}
		if g.ball.x+g.ball.size >= g.bricks[i].x &&
			g.ball.x <= g.bricks[i].x+g.bricks[i].width &&
			g.ball.y+g.ball.size >= g.bricks[i].y &&
			g.ball.y <= g.bricks[i].y+g.bricks[i].height {
			g.bricks[i].active = false
			g.ball.dy = -g.ball.dy
			g.score += g.bricks[i].score
		}
	}

	// Check for game over
	if g.ball.y > screenHeight {
		g.gameOver = true
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw paddle
	ebitenutil.DrawRect(screen, g.paddle.x, g.paddle.y, g.paddle.width, g.paddle.height, color.White)

	// Draw ball
	ebitenutil.DrawRect(screen, g.ball.x, g.ball.y, g.ball.size, g.ball.size, color.White)

	// Draw bricks
	for _, brick := range g.bricks {
		if brick.active {
			ebitenutil.DrawRect(screen, brick.x, brick.y, brick.width, brick.height, brick.color)
		}
	}

	// Draw score and game over text
	if g.gameOver {
		ebitenutil.DebugPrint(screen, "Game Over! Press SPACE to restart")
	}
	ebitenutil.DebugPrint(screen, "Score: "+strconv.Itoa(g.score))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Breakout")

	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
