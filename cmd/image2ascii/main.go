package main

import (
	"flag"
	"fmt"
	"image"
	"io"
	"log"
	"net/http"
	"os"

	_ "golang.org/x/image/webp"
	_ "image/jpeg"
	_ "image/png"

	"github.com/zkck/image2ascii"
)

func sanitizeDimensions(width, height *uint) {
	if *width == 0 && *height == 0 {
		*height = 32
	}
}

func main() {
	var (
		width  = flag.Uint("w", 0, "width of the ascii output")
		height = flag.Uint("h", 0, "height of the ascii output")
		isUrl  = flag.Bool("u", false, "fetch image from URL (instead of file)")
		color  = flag.Bool("c", false, "color the ascii output")
		bold   = flag.Bool("b", false, "bold the ascii output")
	)
	flag.Parse()

	sanitizeDimensions(width, height)

	var reader io.Reader
	pathOrUrl := flag.Arg(0)
	if pathOrUrl == "" {
		log.Fatal(fmt.Errorf("path is required"))
	}

	if *isUrl {
		resp, err := http.Get(pathOrUrl)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		reader = resp.Body
	} else {
		f, err := os.Open(pathOrUrl)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		reader = f
	}

	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	converter := image2ascii.DefaultConverter()
	converter.Color = *color
	converter.Bold = *bold

	ascii := converter.Convert(img, *width, *height)

	fmt.Print(ascii)
}
