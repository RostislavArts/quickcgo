package main

import (
	"fmt"

	"github.com/RostislavArts/quickcgo/quickcg"
)

func main() {
	s, err := quickcg.NewScreen(500, 500, false, "Gradient")
	if err != nil {
		fmt.Println(err)
		return
	}

	color := quickcg.ColorHSV{H: 0, S: 1, V: 1}

	for !quickcg.Done(5) {
		if quickcg.KeyDown(quickcg.KEY_ESCAPE) {
			return
		}

		err := s.Fill(quickcg.HSVtoRGB(color))
		if err != nil {
			fmt.Println(err)
			return
		}
		color.H += 0.001

		s.Redraw()
	}
}

