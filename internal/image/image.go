package image

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/image/draw"
)

const (
	HALF_BLOCK = "▀"         // upper half block character
	RESETLINE  = "\033[0m\n" // select graphic rendition reset + newline
)

// Set the foreground color
func foreground(r, g, b uint32) string {
	return "\033[38;2;" +
		strconv.FormatUint(uint64(r>>8), 10) + ";" +
		strconv.FormatUint(uint64(g>>8), 10) + ";" +
		strconv.FormatUint(uint64(b>>8), 10) + "m"
}

// Set the background color
func background(r, g, b uint32) string {
	return "\033[48;2;" +
		strconv.FormatUint(uint64(r>>8), 10) + ";" +
		strconv.FormatUint(uint64(g>>8), 10) + ";" +
		strconv.FormatUint(uint64(b>>8), 10) + "m"
}

// Draw the album art given a url
func DrawImage(src image.Image, width, height int) string {
	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.CatmullRom.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)
	//draw.BiLinear.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

	b := dst.Bounds()
	var builder strings.Builder

	// Maximum size assumes:
	// 38 characters to set color
	// 1 half block character
	// 5 characters per line for reset+newline
	size := (b.Dx()*39 + 5) * (b.Dy() / 2)

	builder.Grow(size)
	for y := b.Min.Y; y < b.Max.Y; y += 2 {
		for x := b.Min.X; x < b.Max.X; x++ {
			r, g, b, _ := dst.At(x, y).RGBA()
			builder.WriteString(foreground(r, g, b))
			r, g, b, _ = dst.At(x, y+1).RGBA()
			builder.WriteString(background(r, g, b) + HALF_BLOCK)
		}
		builder.WriteString(RESETLINE)
	}

	s := builder.String()
	return s[:len(s)-1] // remove last newline
}

// Get an image by url
func GetImage(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("get image: %v", err)
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("decode image: %v", err)
	}

	return img, nil
}
