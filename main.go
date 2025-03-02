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
	ballSize     = 8
	ballSpeed    = 3
	ballStartY   = screenHeight - 60
)

type Game struct {
	paddle      paddle
	ball        ball
	bricks      []brick
	score       int
	gameOver    bool
	initialized bool
}

type ball struct {
	x, y, dx, dy, size float64
}

func (g *Game) init() {
	// Initialize paddle
	g.paddle = newPaddle()

	// Initialize ball
	g.ball = ball{
		x:    screenWidth / 2,
		y:    ballStartY,
		dx:   ballSpeed,
		dy:   -ballSpeed,
		size: ballSize,
	}

	// Initialize bricks
	g.bricks = initBricks()
	g.initialized = true
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
		g.paddle.moveLeft()
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.paddle.moveRight()
	}

	// Keep paddle within screen bounds
	g.paddle.keepWithinBounds()

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
