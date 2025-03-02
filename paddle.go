package main

const (
	paddleHeight = 12
	paddleWidth  = 60
	paddleSpeed  = 5
	paddleY      = screenHeight - 40
)

func newPaddle() paddle {
	return paddle{
		x:      screenWidth/2 - paddleWidth/2,
		y:      paddleY,
		width:  paddleWidth,
		height: paddleHeight,
	}
}

type paddle struct {
	x, y, width, height float32
}

func (p *paddle) moveLeft() {
	p.x -= paddleSpeed
}

func (p *paddle) moveRight() {
	p.x += paddleSpeed
}

func (p *paddle) keepWithinBounds() {
	if p.x < 0 {
		p.x = 0
	}
	if p.x > screenWidth-p.width {
		p.x = screenWidth - p.width
	}
}
