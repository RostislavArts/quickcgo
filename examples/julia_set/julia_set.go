package main

import (
	"fmt"
	"math/cmplx"

	"github.com/RostislavArts/quickcgo/quickcg"
)

func main() {
	s, err := quickcg.NewScreen(800, 800, false, "Julia Set")
	if err != nil {
		fmt.Println(err)
		return
	}

	c := complex(-0.7, 0.27015)
	zoom := 1.0
	offsetX := 0.0
	offsetY := 0.0
	maxIter := 128

	rendered := false
	for !quickcg.Done(64) {
		if quickcg.KeyDown(quickcg.KEY_ESCAPE) {
			return
		}

		if !rendered {
			for x := range s.GetWidth() {
				for y := range s.GetHeight() {
					scale := 1.5 / zoom

					zx := (float64(x) / float64(s.GetWidth()) - 0.5) * 2 * scale + offsetX
					zy := (float64(y) / float64(s.GetHeight()) - 0.5) * 2 * scale + offsetY
					z := complex(zx, zy)

					var i int
					for i = range maxIter {
						z = z * z + c
						if cmplx.Abs(z) > 2 {
							break
						}
					}

					if i < maxIter {
						t := float64(i) / float64(maxIter)
						r := uint8(9 * (1 - t) * t * t * t * 255)
						g := uint8(15 * (1 - t) * (1 - t) * t * t * 255)
						b := uint8(8.5 * (1 - t) * (1 - t) * (1 - t) * t * 255)

						color := quickcg.ColorRGB{R: r, G: g, B: b}
						s.PSet(x, y, color)
					}
				}
			}

			rendered = true
		}

		quickcg.GetMouseState()
		if quickcg.LMB {
			mouseX := quickcg.MouseX
			mouseY := quickcg.MouseY

			scale := 1.5 / zoom
			oldX := (float64(mouseX) / float64(s.GetWidth()) - 0.5) * 2 * scale + offsetX
			oldY := (float64(mouseY) / float64(s.GetHeight()) - 0.5) * 2 * scale + offsetY

			zoom *= 1.80

			newScale := 1.5 / zoom

			offsetX = oldX - ((float64(mouseX) / float64(s.GetWidth())) - 0.5) * 2 * newScale
			offsetY = oldY - ((float64(mouseY) / float64(s.GetHeight())) - 0.5) * 2 * newScale

			rendered = false
		}
		s.Redraw()
	}
}

