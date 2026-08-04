package sdl3go

/*
#include <stdlib.h>
#include <SDL3/SDL.h>

static bool sdl3go_window_hdr_enabled(SDL_PropertiesID props) {
    return SDL_GetBooleanProperty(props, SDL_PROP_WINDOW_HDR_ENABLED_BOOLEAN, false);
}
static float sdl3go_window_sdr_white_level(SDL_PropertiesID props) {
    return SDL_GetFloatProperty(props, SDL_PROP_WINDOW_SDR_WHITE_LEVEL_FLOAT, 1.0f);
}
static float sdl3go_window_hdr_headroom(SDL_PropertiesID props) {
    return SDL_GetFloatProperty(props, SDL_PROP_WINDOW_HDR_HEADROOM_FLOAT, 1.0f);
}
*/
import "C"
import "unsafe"

type WindowHDRState struct {
	Enabled       bool
	SDRWhiteLevel float32
	Headroom      float32
}

// HDRState reports SDL's current desktop compositor state for this window.
// SDRWhiteLevel is relative to 80 nits and Headroom is display peak divided by
// that white level.
func (w *Window) HDRState() WindowHDRState {
	if w == nil || w.handle == nil {
		return WindowHDRState{SDRWhiteLevel: 1, Headroom: 1}
	}
	properties := C.SDL_GetWindowProperties(w.handle)
	return WindowHDRState{
		Enabled:       bool(C.sdl3go_window_hdr_enabled(properties)),
		SDRWhiteLevel: float32(C.sdl3go_window_sdr_white_level(properties)),
		Headroom:      float32(C.sdl3go_window_hdr_headroom(properties)),
	}
}

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
