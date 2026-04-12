package identicon

import (
	"bytes"
	"image"
	"image/png"

	"github.com/ok-yyyy/go-identicon/internal/digest"
	"github.com/ok-yyyy/go-identicon/internal/palette"
	"github.com/ok-yyyy/go-identicon/internal/pattern"
	"github.com/ok-yyyy/go-identicon/internal/render"
)

const (
	cellSize   = 70
	marginSize = cellSize / 2
)

// Generate creates an identicon image based on the input string.
func Generate(input string) (image.Image, error) {
	sum := digest.Sum([]byte(input))

	grid := pattern.FromDigest(sum)
	fg := palette.Foreground(sum)
	bg := palette.Background()

	img := render.Draw(grid, fg, bg, render.Config{
		CellSize: cellSize,
		Margin:   marginSize,
	})

	return img, nil
}

// EncodePNG generates an identicon image and encodes it as PNG bytes.
func EncodePNG(input string) ([]byte, error) {
	img, err := Generate(input)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
