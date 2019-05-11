package main

import (
	"flag"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"strings"
	"unicode"

	"golang.org/x/image/colornames"
)

var (
	x         = 0
	y         = 0
	height    = flag.Int("h", 100, "Assign the image's height")
	width     = flag.Int("w", 100, "Assign the image's width")
	colorname = flag.String("c", "Red", "Colorize the image")
	filename  = flag.String("n", "sampleImage.jpg", "Name the image")
	quality   = 100
)

func main() {
	log.Println("create start")
	flag.Parse()

	colorRGBA := getColorRGBA(*colorname)
	if colorRGBA == nil {
		log.Println("Cannot find the color. Please kindly confirm it.")
		os.Exit(1)
	}

	img := image.NewRGBA(image.Rect(x, y, *width, *height))

	for i := img.Rect.Min.Y; i < img.Rect.Max.Y; i++ {
		for j := img.Rect.Min.X; j < img.Rect.Max.X; j++ {
			img.Set(j, i, colorRGBA)
		}
	}

	if !strings.HasSuffix(*filename, ".jpg") {
		extension := ".jpg"
		*filename += extension
	}
	file, err := os.Create(*filename)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	err = jpeg.Encode(file, img, &jpeg.Options{quality})
	if err != nil {
		log.Println(err)
	}

	log.Println("create end")
}

func getColorRGBA(colorname string) (c color.Color) {
	if isFirstUpper(colorname) {
		colorname = strings.ToLower(string(rune(colorname[0]))) + colorname[1:]
	}
	colorNameMap := colornames.Map
	for name, rgba := range colorNameMap {
		if colorname == name {
			c = rgba
		}
	}
	return
}

func isFirstUpper(s string) bool {
	r := rune(s[0])
	return unicode.IsUpper(r)
}
