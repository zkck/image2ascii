package ansicodes

import (
	"fmt"
	"image/color"
)

const Reset = "\u001b[0m"

func SetForegroundColor(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("\u001b[38;2;%d;%d;%dm", r>>8, g>>8, b>>8)
}
