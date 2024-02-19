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
	isUrl := flag.Bool("u", false, "path is URL")
	path := flag.String("p", "", "path")
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

	ascii, err := image2ascii.NewConverter(image2ascii.DefaultConfig()).Convert(img)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(ascii)
}
