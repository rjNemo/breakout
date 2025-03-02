package main

import (
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	paddle      paddle
	ball        ball
	bricks      []brick
	score       int
	gameOver    bool
	initialized bool
}

func (g *Game) init() {
	// Initialize paddle
	g.paddle = newPaddle()

	// Initialize ball
	g.ball = newBall()

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
	g.ball.updatePosition()

	// Ball collision with walls
	g.ball.checkWallCollision()

	// Ball collision with paddle
	g.ball.checkPaddleCollision(g.paddle)

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
	vector.DrawFilledRect(screen, g.paddle.x, g.paddle.y, g.paddle.width, g.paddle.height, color.White, true)

	// Draw ball
	vector.DrawFilledRect(screen, g.ball.x, g.ball.y, g.ball.size, g.ball.size, color.White, true)

	// Draw bricks
	for _, brick := range g.bricks {
		if brick.active {
			vector.DrawFilledRect(screen, brick.x, brick.y, brick.width, brick.height, brick.color, true)
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
