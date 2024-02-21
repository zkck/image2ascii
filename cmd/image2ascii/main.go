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

func main() {
	var (
		width  = flag.Uint("w", 0, "width of ascii image")
		height = flag.Uint("h", 32, "height of ascii image")
		isUrl  = flag.Bool("u", false, "path is URL")
		color  = flag.Bool("c", false, "convert to color image")
		bold   = flag.Bool("b", false, "bold")
	)
	flag.Parse()

	var reader io.Reader
	pathOrUrl := flag.Args()[0]

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

	ascii, err := converter.Convert(img, *width, *height)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(ascii)
}
