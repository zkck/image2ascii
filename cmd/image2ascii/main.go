package main

import (
	"flag"
	"fmt"
	"image"
	"io"
	"log"
	"net/http"
	"os"

	_ "image/jpeg"
	_ "image/png"

	"github.com/zkck/image2ascii"
)

func main() {
	var (
		isUrl   = flag.Bool("u", false, "path is URL")
		path    = flag.String("p", "", "path")
		width   = flag.Uint("w", 0, "width")
		height  = flag.Uint("h", 32, "height")
		noColor = flag.Bool("nc", false, "no color")
	)
	flag.Parse()

	var reader io.Reader

	if *isUrl {
		resp, err := http.Get(*path)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		reader = resp.Body
	} else {
		f, err := os.Open(*path)
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
	converter.Color = !(*noColor)

	ascii, err := converter.Convert(img, *width, *height)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(ascii)
}
