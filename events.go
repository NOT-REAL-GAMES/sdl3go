// events.go
package sdl3go

/*
#include <SDL3/SDL.h>

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

	// Pen events
	PenProximity *PenProximityEvent
	PenMotion    *PenMotionEvent
	PenTouch     *PenTouchEvent
	PenButton    *PenButtonEvent
	PenAxis      *PenAxisEvent
}

func PollEvent() (*Event, bool) {
	var cevent C.SDL_Event

	if !C.SDL_PollEvent(&cevent) {
		return nil, false
	}

	event := &Event{
		Type: EventType(C.get_event_type(&cevent)),
	}

	switch event.Type {
	case EVENT_PEN_PROXIMITY_IN, EVENT_PEN_PROXIMITY_OUT:
		event.PenProximity = parsePenProximityEvent(&cevent)
	case EVENT_PEN_MOTION:
		event.PenMotion = parsePenMotionEvent(&cevent)
	case EVENT_PEN_DOWN, EVENT_PEN_UP:
		event.PenTouch = parsePenTouchEvent(&cevent)
	case EVENT_PEN_BUTTON_DOWN, EVENT_PEN_BUTTON_UP:
		event.PenButton = parsePenButtonEvent(&cevent)
	case EVENT_PEN_AXIS:
		event.PenAxis = parsePenAxisEvent(&cevent)
	}

	return event, true
}
