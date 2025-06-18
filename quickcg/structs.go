// QuickCGO is a minimal graphics library in Go inspired by QuickCG (C++) using SDL2.
// It provides pixel drawing, image loading/saving, simple text rendering, and color manipulation.

package quickcg

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Key represents a physical keyboard key using SDL scancodes.
// It is used in functions like KeyDown and IsKeyPressed for input handling.
type Key uint32

// ColorRGB represents a color in the RGB model with 8-bit red, green, and blue channels.
type ColorRGB struct {
	R, G, B uint8
}

// ColorHSL represents a color in the HSL color model (Hue, Saturation, Lightness).
type ColorHSL struct {
	H, S, L float64 // range [0,1]
}

// ColorHSV represents a color in the HSV color model (Hue, Saturation, Value).
type ColorHSV struct {
	H, S, V float64 // range [0,1]
}

// Screen represents an SDL-based window and rendering context.
// It encapsulates the state required to draw, handle events, and interact with a single window.
type Screen struct {
	surface        *sdl.Surface  // the screen surface (legacy; rarely used in SDL2 with renderer)
	window         *sdl.Window   // SDL window object
	windowID	   uint32        // unique SDL window ID for event filtering
	renderer       *sdl.Renderer // SDL renderer for accelerated drawing
	event          sdl.Event     // last received SDL event
	w, h           int           // window width and height in pixels
	buffer []ColorRGB            // logic pixel buffer
	texture *sdl.Texture         // SDL-texture for output
}

var (
	KeyState       []uint8 // keyboard state
	MouseX, MouseY int32   // mouse position
	LMB, RMB       bool    // left and right mouse buttons
)

