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

### SDL2 Note

You do not need to install SDL2 manually — Go bindings are provided via CGO and will link dynamically to the SDL2 shared library at runtime.

However, your system must still have the SDL2 runtime library installed:

On __Ubuntu 22.04 and above__, type:\
`apt install libsdl2{,-image,-mixer,-ttf,-gfx}-dev`

On __Fedora 36 and above__, type:\
`dnf install SDL2{,_image,_mixer,_ttf,_gfx}-devel`

On __Arch Linux__, type:\
`pacman -S sdl2{,_image,_mixer,_ttf,_gfx}`

On __Gentoo__, type:\
`emerge -av libsdl2 sdl2-{image,mixer,ttf,gfx}`

On __macOS__, install SDL2 via [Homebrew](http://brew.sh) like so:\
`brew install sdl2{,_image,_mixer,_ttf,_gfx} pkg-config`

On __Windows__:
Download the SDL2 runtime DLL and place it next to your binary from SDL releases. Or use mingw64 with SDL2 (you can install it from https://github.com/libsdl-org/SDL/releases).

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

## Example: Per-Pixel Drawing (Buffered)

```go
package main

import (
	"github.com/RostislavArts/quickcgo/quickcg"
)

func main() {
	scr, _ := quickcg.NewScreen(256, 256, false, "Fast Drawing Example")

	// Set pixels in internal buffer
	for x := 0; x < 256; x++ {
		for y := 0; y < 256; y++ {
			scr.WritePixel(x, y, quickcg.ColorRGB{R: uint8(x), G: uint8(y), B: 128})
		}
	}

	// Draw buffer to screen
	scr.DrawBuffer()

	scr.Redraw()
	scr.Sleep()
}
```

This is **much faster** than calling `PSet()` per pixel, especially in real-time rendering.
For immediate pixel updates, you can still use `PSet`.

You can check more examples in [*examples*](https://github.com/RostislavArts/quickcgo/tree/main/examples) folder.

## Library Structure

### Types

- `Screen` — represents a rendering window
- `ColorRGB`, `ColorHSL`, `ColorHSV` — color types

### Core Methods

- `NewScreen(width, height, fullscreen, title) *Screen`
- `(*Screen).PSet(x, y int, color ColorRGB)` — set a pixel
- `(*Screen).Redraw()` — update the screen
- `(*Screen).WritePixel(x, y int, color ColorRGB)` — write pixel to buffer (fast)
- `(*Screen).DrawBuffer()` — update screen from buffer
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

## Performance Notes

* Prefer `WritePixel()` + `DrawBuffer()` when drawing many pixels.
* Avoid calling `PSet()` inside loops — it’s slow and not hardware-accelerated.
* Buffered drawing uses a texture and streams the entire frame to the GPU once per frame.

## Concurrency & Thread Safety

SDL is **not thread-safe**, particularly `sdl.PollEvent()`. Using few windows at the same time may cause problems.

## License

MIT License

---

### Based On

- [QuickCG (C++)](https://lodev.org/quickcg/)
- [go-sdl2 bindings](https://github.com/veandco/go-sdl2)
