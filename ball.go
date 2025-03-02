package main

import "math/rand"

const (
	ballSize   = 8
	ballSpeed  = 3
	ballStartY = screenHeight - 60
)

type ball struct {
	x, y, dx, dy, size float32
}

func newBall() ball {
	return ball{
		x:    screenWidth / 2,
		y:    ballStartY,
		dx:   ballSpeed,
		dy:   -ballSpeed,
		size: ballSize,
	}
}

func (b *ball) updatePosition() {
	b.x += b.dx
	b.y += b.dy
}

func (b *ball) checkWallCollision() {
	if b.x <= 0 || b.x >= screenWidth-b.size {
		b.dx = -b.dx
	}
	if b.y <= 0 {
		b.dy = -b.dy
	}
}

func (b *ball) checkPaddleCollision(p paddle) {
	if b.y+b.size >= p.y &&
		b.x+b.size >= p.x &&
		b.x <= p.x+p.width &&
		b.y <= p.y+p.height {
		b.dy = -b.dy
		// Add some randomness to the ball direction
		b.dx += (rand.Float32()*2 - 1) * 0.5
	}
}
