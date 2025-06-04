package quickcg

import (
	"github.com/veandco/go-sdl2/sdl"
)

var previousKeyState []uint8

func readKeys() {
	sdl.PumpEvents()
	KeyState = sdl.GetKeyboardState()
}

// KeyDown returns true if the specified key is currently held down.
func KeyDown(keycode Key) bool {
	readKeys()
	return KeyState[keycode] != 0
}

// KeyPressed returns true if the key was just pressed (edge-triggered).
func KeyPressed(keycode Key) bool {
	readKeys()
	down := KeyState[keycode] != 0 && (previousKeyState == nil || previousKeyState[keycode] == 0)
	previousKeyState = append([]uint8(nil), KeyState...)
	return down
}

// GetMouseState updates mouse position and button states.
func GetMouseState() {
	var mouseState uint32

	MouseX, MouseY, mouseState = sdl.GetMouseState()
	LMB = (mouseState & sdl.ButtonLMask()) != 0
	RMB = (mouseState & sdl.ButtonRMask()) != 0
}
