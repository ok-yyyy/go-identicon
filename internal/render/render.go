package render

import (
	"image"
	"image/color"
	"image/draw"
)

const (
	gridSize = 5
)

type Config struct {
	CellSize int
	Margin   int
}

// Draw creates an identicon image based on the provided grid and colors.
func Draw(grid [5][5]bool, fg, bg color.RGBA, cfg Config) *image.RGBA {
	if cfg.CellSize <= 0 {
		cfg.CellSize = 32
	}
	if cfg.Margin < 0 {
		cfg.Margin = 0
	}

	width := cfg.CellSize*gridSize + 2*cfg.Margin
	rect := image.Rect(0, 0, width, width)
	img := image.NewRGBA(rect)

	draw.Draw(img, rect, &image.Uniform{C: bg}, image.Point{}, draw.Src)
	drawGrid(img, grid, fg, cfg)

	return img
}

func drawGrid(img *image.RGBA, grid [5][5]bool, fg color.RGBA, cfg Config) {
	fgImg := &image.Uniform{fg}

	for row := 0; row < gridSize; row++ {
		for col := 0; col < gridSize; col++ {
			if !grid[row][col] {
				continue
			}

			x0 := cfg.Margin + col*cfg.CellSize
			y0 := cfg.Margin + row*cfg.CellSize
			cell := image.Rect(x0, y0, x0+cfg.CellSize, y0+cfg.CellSize)
			draw.Draw(img, cell, fgImg, image.Point{}, draw.Src)
		}
	}
}
