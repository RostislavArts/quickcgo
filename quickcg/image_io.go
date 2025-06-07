package quickcg

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"time"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

// LoadPNG loads a PNG image from directory and returns its pixels (row-major), width and height.
func LoadPNG(path string) ([]ColorRGB, int, int, error) {
	file, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("Error loading file: %s", err)
		return nil, 0, 0, err
	}

	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		err = fmt.Errorf("Error decoding file: %s", err)
		return nil, 0, 0, err
	}

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	pixels := make([]ColorRGB, width*height)

	for y := range height {
		for x := range width {
			r, g, b, _ := img.At(x, y).RGBA()
			pixels[y*width+x] = ColorRGB{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
			}
		}
	}

	return pixels, width, height, nil
}

// SaveScreenAsPNG saves the current contents of the screen to a PNG file at the given path.
func (scr *Screen) SaveScreenAsPNG(path string) error {
	img := image.NewRGBA(image.Rect(0, 0, scr.w, scr.h))

	pixelData := make([]byte, scr.w*scr.h*4)

	err := scr.renderer.ReadPixels(
		nil,
		sdl.PIXELFORMAT_ABGR8888,
		unsafe.Pointer(&pixelData[0]),
		scr.w*4,
	)
	if err != nil {
		return fmt.Errorf("failed to read pixels: %w", err)
	}

	for y := range scr.h {
		for x := range scr.w {
			i := (y*scr.w + x) * 4
			r := pixelData[i]
			g := pixelData[i+1]
			b := pixelData[i+2]
			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}

	path = fmt.Sprintf("%v/screenshot_%v.png", path, time.Now().Format("2006-01-02_15-04-05"))
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return fmt.Errorf("failed to encode PNG: %w", err)
	}

	return nil
}
