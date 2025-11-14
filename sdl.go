package sdl3go

/*
#cgo windows LDFLAGS: -lSDL3
#cgo linux LDFLAGS: -lSDL3
#cgo darwin LDFLAGS: -lSDL3
#include <SDL3/SDL.h>
*/
import "C"

type InitFlags uint32

const (
	INIT_VIDEO InitFlags = C.SDL_INIT_VIDEO
)

func Init(flags InitFlags) error {
	if !C.SDL_Init(C.Uint32(flags)) {
		return GetError()
	}
	return nil
}

func Quit() {
	C.SDL_Quit()
}

func GetError() error {
	errStr := C.GoString(C.SDL_GetError())
	if errStr == "" {
		return nil
	}
	return &Error{Message: errStr}
}

type Error struct {
	Message string
}

func (e *Error) Error() string {
	return e.Message
}
