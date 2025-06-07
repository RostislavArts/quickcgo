package main

import (
	"fmt"
	"sync"
	"math"

	"github.com/RostislavArts/quickcgo/quickcg"
)

func main() {
	var wg sync.WaitGroup

	s1, err := quickcg.NewScreen(256, 256, false, "Window 1")
	if err != nil {
		fmt.Println(err)
		return
	}

	s2, err := quickcg.NewScreen(256, 256, false, "Window 2")
	if err != nil {
		fmt.Println(err)
		return
	}

	for x := range s1.GetWidth()  {
		for y := range s1.GetHeight()  {
			err := s1.PSet(x, y, quickcg.ColorRGB{R: uint8(math.Sin(float64(x))), G: uint8(y), B: 128})
			if err != nil {
				print(err)
				return
			}
		}
	}

	for x := range s2.GetWidth()  {
		for y := range s2.GetHeight()  {
			err := s2.PSet(y, x, quickcg.ColorRGB{R: uint8(x), G: uint8(math.Cos(float64(y))), B: 128})
			if err != nil {
				print(err)
				return
			}
		}
	}

	s1.Redraw()
	s2.Redraw()

	wg.Add(2)
	go sleep(s1, &wg)
	go sleep(s2, &wg)

	wg.Wait()
}

func sleep(scr *quickcg.Screen, wg *sync.WaitGroup) {
	defer wg.Done()
	scr.Sleep()
}
