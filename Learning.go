package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

const (
	width      = 1080 // IG post size
	height     = 1350 // 4:5 ratio
	dropCount  = 800
	dropLength = 15
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 1. Create canvas with dark rainy sky
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	darkBlue := color.RGBA{25, 35, 60, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{darkBlue}, image.Point{}, draw.Src)

	// 2. Draw rain using goroutines - intermediate flex
	dropChan := make(chan image.Point, dropCount)
	for i := 0; i < 4; i++ { // 4 workers
		go func() {
			for pos := range dropChan {
				drawRainDrop(img, pos.X, pos.Y)
			}
		}()
	}

	for i := 0; i < dropCount; i++ {
		x := rand.Intn(width)
		y := rand.Intn(height)
		dropChan <- image.Point{x, y}
	}
	close(dropChan)
	time.Sleep(100 * time.Millisecond) // let goroutines finish

	// 3. Draw "GO" text like it's written on foggy glass
	addLabel(img, 100, 300, "Rainy Day", color.RGBA{200, 200, 220, 255})
	addLabel(img, 100, 400, "with my GO", color.RGBA{0, 173, 216, 255}) // Go cyan
	addLabel(img, 100, 500, "// stay dry, keep coding", color.RGBA{150, 150, 160, 255})

	// 4. Save it
	f, err := os.Create("rainy_day.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}

func drawRainDrop(img *image.RGBA, x, y int) {
	rainColor := color.RGBA{180, 200, 255, 180} // translucent
	for i := 0; i < dropLength; i++ {
		if y+i < height {
			img.Set(x, y+i, rainColor)
		}
	}
}

func addLabel(img *image.RGBA, x, y int, label string, col color.RGBA) {
	point := fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}
