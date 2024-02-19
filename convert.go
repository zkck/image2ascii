package image2ascii

import (
	"fmt"
	"image"
	"image/color"
	"strings"
)

const ansiResetCode = "\u001b[0m"

type Converter struct {
	AsciiMap string
	Color    bool
}

func DefaultConverter() Converter {
	return Converter{
		AsciiMap: " .:-=+*#%@",
		Color:    true,
	}
}

func getAnsiColorCode(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("\u001b[38;2;%d;%d;%dm", r&0xff, g&0xff, b&0xff)
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
			pixelDepth, _, _, _ := color.GrayModel.Convert(img.At(x, y)).RGBA()
			asciiChar := c.AsciiMap[len(c.AsciiMap)*int(pixelDepth)/(0xffff+1)]
			if c.Color {
				builder.WriteString(getAnsiColorCode(img.At(x, y)))
				builder.WriteByte(asciiChar)
				builder.WriteString(ansiResetCode)
			} else {
				builder.WriteByte(asciiChar)
			}
		}
		builder.WriteByte('\n')
	}

	return builder.String(), nil
}
