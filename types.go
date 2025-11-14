package sdl3go

/*
#include <SDL3/SDL.h>
*/
import "C"

type Window struct {
	handle *C.SDL_Window
}

type WindowFlags uint32

const (
	WINDOW_VULKAN WindowFlags = C.SDL_WINDOW_VULKAN
)
