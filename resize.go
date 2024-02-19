package image2ascii

import (
	"fmt"
	"image"

	"golang.org/x/image/draw"
)

func scaleBounds(bounds image.Rectangle, width, height uint) (image.Rectangle, error) {
	var p image.Point
	if width == 0 && height == 0 {
		return image.Rectangle{}, fmt.Errorf("either width or height must be non-zero")
	} else if width == 0 {
		p.X, p.Y = 2*bounds.Dx()*int(height)/bounds.Dy(), int(height)
	} else if height == 0 {
		p.X, p.Y = int(width), bounds.Dy()*int(width)/(2*bounds.Dx())
	} else {
		p.X, p.Y = int(width), int(height)
	}
	return image.Rectangle{Max: p}, nil
}

func resize(src image.Image, dstBounds image.Rectangle) image.Image {
	dst := image.NewRGBA(dstBounds)
	draw.BiLinear.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Over, nil)
	return dst
}
