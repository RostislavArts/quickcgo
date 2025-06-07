# QuickCGO

**QuickCGO** is a Go port of the [QuickCG](https://lodev.org/quickcg/) C++ graphics library. It provides a simple API for 2D graphics, color manipulation, image I/O, text rendering, and keyboard/mouse input using [SDL2](https://github.com/libsdl-org/SDL). The aim of this project is to make an easy-to-use graphics library that you can use to learn 2D graphics.

## Features

- Window creation and pixel-level rendering
- Drawing primitives: lines, rectangles, circles, filled circles
- Basic text rendering using `golang.org/x/image/font`
- PNG image loading and saving
- Color conversion: RGB ↔ HSL / HSV
- Keyboard and mouse input
- Optional support for multiple windows and concurrent rendering (not very stable)

## Installation

To use `quickcgo`, all you need is **Go 1.22 or later** and Go modules enabled.

In your project, add this import:

```go
import "github.com/RostislavArts/quickcgo/quickcg"
```

Then run:

```bash
go mod tidy
```

This will automatically fetch quickcgo and its dependencies.

## Example

```go
package main

import (
	"github.com/RostislavArts/quickcgo/quickcg"
)

func main() {
	scr, _ := quickcg.NewScreen(256, 256, false, "Hello World")

	for x := 0; x < 256; x++ {
		for y := 0; y < 256; y++ {
			scr.PSet(x, y, quickcg.ColorRGB{R: uint8(x), G: uint8(y), B: 128})
		}
	}

	scr.Redraw()
	scr.Sleep()
}
```

You can check more examples in [*examples*](https://github.com/RostislavArts/quickcgo/tree/main/examples) folder.

## Library Structure

### Types

- `Screen` — represents a rendering window
- `ColorRGB`, `ColorHSL`, `ColorHSV` — color types

### Core Methods

- `NewScreen(width, height, fullscreen, title) *Screen`
- `(*Screen).PSet(x, y int, color ColorRGB)` — set a pixel
- `(*Screen).Redraw()` — update the screen
- `(*Screen).Fill(color ColorRGB)` — fill screen with color
- Drawing:
  - `DrawLine`, `DrawRect`, `DrawCircle`, `DrawFilledCircle`
- Text:
  - `DrawText(x, y int, text string, color ColorRGB)`
- Images:
  - `LoadPNG(path string)`, `SavePNG(path string)`
- Color Conversion:
  - `RGBtoHSL`, `HSLtoRGB`, `RGBtoHSV`, `HSVtoRGB`
- Input:
  - `KeyPressed(keycode)`, `KeyDown(keycode)`
  - `GetMouseState()`, `MouseX`, `MouseY`, `LMB`, `RMB`

## Concurrency & Thread Safety

SDL is **not thread-safe**, particularly `sdl.PollEvent()`. Using few windows at the same time may cause problems.

## License

MIT License

---

### Based On

- [QuickCG (C++)](https://lodev.org/quickcg/)
- [go-sdl2 bindings](https://github.com/veandco/go-sdl2)
