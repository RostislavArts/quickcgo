package main

import (
	"fmt"

	"github.com/RostislavArts/quickcgo/quickcg"
)

func main() {
	scr, err := quickcg.NewScreen(256, 256, false, "Hello World")
	if err != nil {
		fmt.Println(err)
		return
	}

	for x := range scr.GetWidth() {
		for y := range scr.GetHeight() {
			err := scr.PSet(x, y, quickcg.ColorRGB{R: uint8(x), G: uint8(y), B: 128})
			if err != nil {
				print(err)
				return
			}
		}
	}

	scr.Redraw()
	err = scr.Sleep()
	if err != nil {
		fmt.Println(err)
		quickcg.Quit()
		return
	}
}
