package quickcg

import (
	"fmt"
	"image"
	colorLib "image/color"

	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

// PSet sets the pixel at (x, y) to the given RGB color.
func (screen *Screen) PSet(x, y int, color ColorRGB) error {
	if x < 0 || y < 0 || x >= screen.w || y >= screen.h {
		err := fmt.Errorf("Can't place pixel out of window bounds!")
		return err
	}

	var err error

	err = screen.renderer.SetDrawColor(color.R, color.G, color.B, 255)
	if err != nil {
		err = fmt.Errorf("Error setting draw color: %s", err)
		return err
	}

	err = screen.renderer.DrawPoint(int32(x), int32(y))
	if err != nil {
		err = fmt.Errorf("Error drawing pixel: %s", err)
		return err
	}

	return nil
}

// Fill fills the screen with the specified RGB color.
func (screen *Screen) Fill(color ColorRGB) error {
	var err error

	err = screen.renderer.SetDrawColor(color.R, color.G, color.B, 255)
	if err != nil {
		err = fmt.Errorf("Error setting draw color: %s", err)
		return err
	}

	err = screen.renderer.Clear()
	if err != nil {
		err = fmt.Errorf("Error clearing renderer: %s", err)
		return err
	}

	return nil
}

// Redraw updates the display with any changes made since the last call.
func (screen *Screen) Redraw() {
	screen.renderer.Present()
}

// DrawLine draws a line between two points with the specified color.
func (screen *Screen) DrawLine(x1, y1, x2, y2 int, color ColorRGB) error {
	var err error

	err = screen.renderer.SetDrawColor(color.R, color.G, color.B, 255)
	if err != nil {
		err = fmt.Errorf("Error setting draw color: %s", err)
		return err
	}

	err = screen.renderer.DrawLine(int32(x1), int32(y1), int32(x2), int32(y2))
	if err != nil {
		err = fmt.Errorf("Error drawing line: %s", err)
		return err
	}

	return nil
}

// DrawRect draws the outline of a rectangle.
func (screen *Screen) DrawRect(x1, y1, x2, y2 int, color ColorRGB) error {
	rect := sdl.Rect{
		X: int32(x1),
		Y: int32(y1),
		W: int32(x2 - x1),
		H: int32(y2 - y1),
	}
	var err error
	err = screen.renderer.SetDrawColor(color.R, color.G, color.B, 255)
	if err != nil {
		err = fmt.Errorf("Error setting draw color: %s", err)
		return err
	}

	err = screen.renderer.FillRect(&rect)
	if err != nil {
		err = fmt.Errorf("Error filling rectangle: %s", err)
		return err
	}

	return nil
}

// DrawCircle draws the outline of a circle.
func (screen *Screen) DrawCircle(xc, yc, radius int, color ColorRGB) error {
	x := 0
	y := radius
	d := 3 - 2*radius
	for x <= y {
		points := [][2]int{
			{xc + x, yc + y}, {xc - x, yc + y},
			{xc + x, yc - y}, {xc - x, yc - y},
			{xc + y, yc + x}, {xc - y, yc + x},
			{xc + y, yc - x}, {xc - y, yc - x},
		}

		for _, p := range points {
			err := screen.PSet(p[0], p[1], color)
			if err != nil {
				return fmt.Errorf("DrawCircle failed: %s", err)
			}
		}
		if d < 0 {
			d += 4*x + 6
		} else {
			d += 4*(x - y) + 10
			y--
		}
		x++
	}
	return nil
}

// DrawFilledCircle draws a filled circle centered at (xc, yc) with radius r.
func (screen *Screen) DrawFilledCircle(xc, yc, radius int, color ColorRGB) error {
	for y := -radius; y <= radius; y++ {
		for x := -radius; x <= radius; x++ {
			if x*x+y*y <= radius*radius {
				err := screen.PSet(xc+x, yc+y, color)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// DrawText draws a string on the screen at the specified (x, y) coordinates using the given color.
func (screen *Screen) DrawText(x, y int, text string, color ColorRGB) error {
	col := colorLib.RGBA{R: color.R, G: color.G, B: color.B, A: 255}
	img := image.NewRGBA(image.Rect(0, 0, screen.w, screen.h))
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)},
	}
	d.DrawString(text)
	for py := range screen.h {
		for px := range screen.w {
			clr := img.At(px, py)
			r, g, b, a := clr.RGBA()
			if a > 0 {
				err := screen.PSet(px, py, ColorRGB{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)})
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// DrawImage draws a preloaded image pixel buffer at the given screen position.
func (screen *Screen) DrawImage(pixels []ColorRGB, imgW, imgH, posX, posY int) error {
	for y := range imgH {
		for x := range imgW {
			color := pixels[y*imgW+x]
			err := screen.PSet(posX+x, posY+y, color)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

