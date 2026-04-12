package palette

import (
	"image/color"
	"math"
)

// Background returns a background color.
func Background() color.RGBA {
	return color.RGBA{0xf0, 0xf0, 0xf0, 0xff}
}

// Foreground returns a foreground color based on the input.
func Foreground(data [16]byte) color.RGBA {
	remap := func(v, vmin, vmax int, dmin, dmax float64) float64 {
		if v <= vmin {
			return dmin
		}
		if v >= vmax {
			return dmax
		}
		return float64(v-vmin)/float64(vmax-vmin)*(dmax-dmin) + dmin
	}

	h := int(data[12]&0x0f)<<8 | int(data[13])
	s := int(data[14])
	l := int(data[15])

	hue := remap(h, 0, 4095, 0, 360)
	sat := remap(s, 0, 255, 0, 20)
	lum := remap(l, 0, 255, 0, 20)

	return hslToRGB(hue, 65-sat, 75-lum)
}

func hslToRGB(h, s, l float64) color.RGBA {
	var a, b float64

	hue := h / 360.0
	sat := s / 100.0
	lum := l / 100.0

	if lum < 0.5 {
		b = lum * (sat + 1.0)
	} else {
		b = lum + sat - lum*sat
	}
	a = lum*2.0 - b

	return color.RGBA{
		uint8(math.Round(255 * hueToRGB(a, b, hue+1.0/3.0))),
		uint8(math.Round(255 * hueToRGB(a, b, hue))),
		uint8(math.Round(255 * hueToRGB(a, b, hue-1.0/3.0))),
		255,
	}
}

func hueToRGB(a, b, hue float64) float64 {
	h := hue
	if hue < 0.0 {
		h = hue + 1.0
	} else if hue > 1.0 {
		h = hue - 1.0
	}

	if h < 1.0/6.0 {
		return a + (b-a)*6.0*h
	}
	if h < 1.0/2.0 {
		return b
	}
	if h < 2.0/3.0 {
		return a + (b-a)*(2.0/3.0-h)*6.0
	}
	return a
}
