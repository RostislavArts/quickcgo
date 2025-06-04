package quickcg

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
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
func (screen *Screen) SaveScreenAsPNG(path string) error {
	img := image.NewRGBA(image.Rect(0, 0, screen.w, screen.h))
	surface, err := screen.window.GetSurface()
	if err != nil {
		err = fmt.Errorf("Error getting surface: %s", err)
		return err
	}

	err = surface.Lock()
	if err != nil {
		err = fmt.Errorf("Error locking surface: %s", err)
		return err
	}

	defer surface.Unlock()
	pixels := surface.Pixels()

	for y := range screen.h {
		for x := range screen.w {
			offset := (y*int(screen.surface.Pitch) + x*4)
			r := pixels[offset+2]
			g := pixels[offset+1]
			b := pixels[offset+0]
			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}

	file, err := os.Create(path)
	if err != nil {
		err = fmt.Errorf("Error creating file: %s", err)
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		err = fmt.Errorf("Error encoding file: %s", err)
		return err
	}

	return nil
}

