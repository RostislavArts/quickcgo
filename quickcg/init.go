package quickcg

import (
	"os"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// NewScreen creates and initializes a new SDL2 window and renderer.
// It takes the window width and height, fullscreen flag, and a title.
// On failure, it logs the error and terminates the program.
// Returns a pointer to a fully initialized Screen struct instance.
func NewScreen(width, height int, fullscreen bool, title string) (*Screen, error) {
	scr := Screen{}

	scr.w = width
	scr.h = height

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		err = fmt.Errorf("Failed to initialize SDL: %s", err)
		return nil, err
	}

	flags := uint32(sdl.WINDOW_SHOWN)
	if fullscreen {
		flags |= sdl.WINDOW_FULLSCREEN
	}

	var err error
	scr.window, err = sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(scr.w), int32(scr.h), flags)
	if err != nil {
		err = fmt.Errorf("Failed to create window: %s", err)
		return nil, err
	}

	scr.windowID, err = scr.window.GetID()
	if err != nil {
		err = fmt.Errorf("Failed to get window id: %s", err)
		return nil, err
	}

	scr.renderer, err = sdl.CreateRenderer(scr.window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		err = fmt.Errorf("Failed to create renderer: %s", err)
		return nil, err
	}

	scr.surface, err = scr.window.GetSurface()
	if err != nil {
		err = fmt.Errorf("Failed to get screen surface: %s", err)
		return nil, err
	}

	return &scr, nil
}

// GetWidth returns the width of the screen in pixels.
func (screen *Screen) GetWidth() int {
	return screen.w
}

// GetHeight returns the height of the screen in pixels.
func (screen *Screen) GetHeight() int {
	return screen.h
}

// Done polls SDL events and returns true if a quit event (e.g. window close)
// has been received. It also inserts an optional delay (in milliseconds) 
// to reduce CPU usage during polling loops.
//
// Use this function in your main loop to detect when the user has requested
// the program to exit (e.g. by clicking the window close button).
//
// The delay parameter controls how long the function sleeps before polling events.
func Done(delay uint32) bool {
	sdl.Delay(delay)

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			return true
		}
	}

	return false
}

// Close safely destroys the window and renderer resources associated with the screen.
// Should be called when the screen is no longer needed.
func (screen *Screen) Close() error {
	var err error

	err = screen.renderer.Destroy()
	if err != nil {
		err = fmt.Errorf("Failed to destroy renderer: %s", err)
		return err
	}

	err = screen.window.Destroy()
	if err != nil {
		err = fmt.Errorf("Failed to destroy window: %s", err)
		return err
	}

	return nil
}

// Sleep blocks execution until the window is closed via the close button (X)
// or a global SDL_QuitEvent is received.
// It polls events with a small delay to avoid high CPU usage.
func (screen *Screen) Sleep() error {
    for {
        for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
            switch e := event.(type) {
            case *sdl.QuitEvent:
				err := screen.Close() 
				if err != nil {
					return err
				}

			 	return nil
			case *sdl.WindowEvent:
				if e.Event == sdl.WINDOWEVENT_CLOSE && e.WindowID == screen.windowID {
					err := screen.Close()
					if err != nil {
						return err
					}

					return nil
				}
            }
        }

        sdl.Delay(5)
    }
}

// Quit shuts down SDL and exits the program.
func Quit() {
 	sdl.Quit()
	os.Exit(0)
}

