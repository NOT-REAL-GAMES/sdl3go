package sdl3go

/*
#include <stdlib.h>
#include <SDL3/SDL.h>
*/
import "C"
import "unsafe"

func CreateWindow(title string, width, height int, flags WindowFlags) (*Window, error) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	window := C.SDL_CreateWindow(cTitle, C.int(width), C.int(height), C.SDL_WindowFlags(flags))
	if window == nil {
		return nil, GetError()
	}

	return &Window{handle: window}, nil
}

func (w *Window) Destroy() {
	C.SDL_DestroyWindow(w.handle)
}
