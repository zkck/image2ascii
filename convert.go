package image2ascii

import (
	"image"
	"image/color"
	"strings"

	"github.com/zkck/image2ascii/ansicodes"
)

type Converter struct {
	AsciiMap string
	Color    bool
	Bold     bool
}

func DefaultConverter() Converter {
	return Converter{
		AsciiMap: " .:-=+*#%@",
		Color:    true,
		Bold:     false,
	}
}

func getColorDepth(c color.Color) int {
	depth, _, _, _ := color.GrayModel.Convert(c).RGBA()
	return int(depth)
}

func (c Converter) Convert(img image.Image, width, height uint) (string, error) {
	if newBounds, err := scaleBounds(img.Bounds(), width, height); err != nil {
		return "", err
	} else {
		img = resize(img, newBounds)
	}

	var builder strings.Builder

	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			asciiChar := c.AsciiMap[len(c.AsciiMap)*getColorDepth(img.At(x, y))/(0xffff+1)]
			if c.Color {
				builder.WriteString(ansicodes.SetForegroundColor(img.At(x, y)))
			}
			if c.Bold {
				builder.WriteString(ansicodes.Bold)
			}
			builder.WriteByte(asciiChar)
			builder.WriteString(ansicodes.Reset)
		}
		builder.WriteByte('\n')
	}

	return builder.String(), nil
}
