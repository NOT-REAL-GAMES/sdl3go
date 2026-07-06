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

func (w *Window) GetSize() (int, int, error) {
	var width C.int
	var height C.int
	if !C.SDL_GetWindowSize(w.handle, &width, &height) {
		return 0, 0, GetError()
	}
	return int(width), int(height), nil
}

func (w *Window) GetSizeInPixels() (int, int, error) {
	var width C.int
	var height C.int
	if !C.SDL_GetWindowSizeInPixels(w.handle, &width, &height) {
		return 0, 0, GetError()
	}
	return int(width), int(height), nil
}
