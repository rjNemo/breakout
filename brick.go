package main

import (
	"image/color"
)

const (
	brickWidth  = 60
	brickHeight = 20
	brickRows   = 5
	brickCols   = 10
	brickGap    = 4
	brickOffset = 40
)

type brick struct {
	x, y, width, height float32
	active              bool
	color               color.Color
	score               int
}

func initBricks() []brick {
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

	bricks := make([]brick, 0, brickRows*brickCols)
	for row := range brickRows {
		for col := range brickCols {
			bricks = append(bricks, brick{
				x:      float32(col*(brickWidth+brickGap) + brickGap),
				y:      float32(row*(brickHeight+brickGap) + brickGap + brickOffset),
				width:  brickWidth,
				height: brickHeight,
				active: true,
				color:  brickConfig[row].color,
				score:  brickConfig[row].score,
			})
		}
	}
	return bricks
}
