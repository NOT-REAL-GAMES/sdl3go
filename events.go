// events.go
package sdl3go

/*
#include <SDL3/SDL.h>

// Helper to get event type since 'type' is a Go keyword
static inline Uint32 get_event_type(SDL_Event *event) {
    return event->type;
}
*/
import "C"

type EventType uint32

const (
	EVENT_QUIT     EventType = C.SDL_EVENT_QUIT
	EVENT_KEY_DOWN EventType = C.SDL_EVENT_KEY_DOWN
	EVENT_KEY_UP   EventType = C.SDL_EVENT_KEY_UP
)

type Event struct {
	Type EventType
}

func PollEvent() (*Event, bool) {
	var cevent C.SDL_Event

	for C.SDL_PollEvent(&cevent) {
		return &Event{
			Type: EventType(C.get_event_type(&cevent)), // Use our helper!
		}, true
	}

	return nil, false
}
