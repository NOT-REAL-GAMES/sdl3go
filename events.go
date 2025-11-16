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
	EVENT_QUIT              EventType = C.SDL_EVENT_QUIT
	EVENT_KEY_DOWN          EventType = C.SDL_EVENT_KEY_DOWN
	EVENT_KEY_UP            EventType = C.SDL_EVENT_KEY_UP
	EVENT_PEN_PROXIMITY_IN  EventType = C.SDL_EVENT_PEN_PROXIMITY_IN
	EVENT_PEN_PROXIMITY_OUT EventType = C.SDL_EVENT_PEN_PROXIMITY_OUT
	EVENT_PEN_DOWN          EventType = C.SDL_EVENT_PEN_DOWN
	EVENT_PEN_UP            EventType = C.SDL_EVENT_PEN_UP
	EVENT_PEN_MOTION        EventType = C.SDL_EVENT_PEN_MOTION
	EVENT_PEN_BUTTON_DOWN   EventType = C.SDL_EVENT_PEN_BUTTON_DOWN
	EVENT_PEN_BUTTON_UP     EventType = C.SDL_EVENT_PEN_BUTTON_UP
	EVENT_PEN_AXIS          EventType = C.SDL_EVENT_PEN_AXIS
	EVENT_DROP_BEGIN        EventType = C.SDL_EVENT_DROP_BEGIN
	EVENT_DROP_FILE         EventType = C.SDL_EVENT_DROP_FILE
	EVENT_DROP_TEXT         EventType = C.SDL_EVENT_DROP_TEXT
	EVENT_DROP_COMPLETE     EventType = C.SDL_EVENT_DROP_COMPLETE
	EVENT_DROP_POSITION     EventType = C.SDL_EVENT_DROP_POSITION
)

type Event struct {
	Type EventType

	// Pen events
	PenProximity *PenProximityEvent
	PenMotion    *PenMotionEvent
	PenTouch     *PenTouchEvent
	PenButton    *PenButtonEvent
	PenAxis      *PenAxisEvent

	// Drag-and-drop events
	Drop *DropEvent
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
	case EVENT_DROP_BEGIN, EVENT_DROP_FILE, EVENT_DROP_TEXT, EVENT_DROP_COMPLETE, EVENT_DROP_POSITION:
		event.Drop = parseDropEvent(&cevent)
	}

	return event, true
}
