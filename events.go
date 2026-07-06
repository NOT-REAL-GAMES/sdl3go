// events.go
package sdl3go

/*
#include <SDL3/SDL.h>

static inline Uint32 get_event_type(SDL_Event *event) {
    return event->type;
}

static inline Uint64 get_window_timestamp(SDL_Event *event) {
    return event->window.timestamp;
}

static inline SDL_WindowID get_window_window_id(SDL_Event *event) {
    return event->window.windowID;
}

static inline Sint32 get_window_data1(SDL_Event *event) {
    return event->window.data1;
}

static inline Sint32 get_window_data2(SDL_Event *event) {
    return event->window.data2;
}
*/
import "C"

type EventType uint32

const (
	EVENT_QUIT                      EventType = C.SDL_EVENT_QUIT
	EVENT_WINDOW_RESIZED            EventType = C.SDL_EVENT_WINDOW_RESIZED
	EVENT_WINDOW_PIXEL_SIZE_CHANGED EventType = C.SDL_EVENT_WINDOW_PIXEL_SIZE_CHANGED
	EVENT_KEY_DOWN                  EventType = C.SDL_EVENT_KEY_DOWN
	EVENT_KEY_UP                    EventType = C.SDL_EVENT_KEY_UP
	EVENT_MOUSE_MOTION              EventType = C.SDL_EVENT_MOUSE_MOTION
	EVENT_MOUSE_BUTTON_DOWN         EventType = C.SDL_EVENT_MOUSE_BUTTON_DOWN
	EVENT_MOUSE_BUTTON_UP           EventType = C.SDL_EVENT_MOUSE_BUTTON_UP
	EVENT_MOUSE_WHEEL               EventType = C.SDL_EVENT_MOUSE_WHEEL
	EVENT_PEN_PROXIMITY_IN          EventType = C.SDL_EVENT_PEN_PROXIMITY_IN
	EVENT_PEN_PROXIMITY_OUT         EventType = C.SDL_EVENT_PEN_PROXIMITY_OUT
	EVENT_PEN_DOWN                  EventType = C.SDL_EVENT_PEN_DOWN
	EVENT_PEN_UP                    EventType = C.SDL_EVENT_PEN_UP
	EVENT_PEN_MOTION                EventType = C.SDL_EVENT_PEN_MOTION
	EVENT_PEN_BUTTON_DOWN           EventType = C.SDL_EVENT_PEN_BUTTON_DOWN
	EVENT_PEN_BUTTON_UP             EventType = C.SDL_EVENT_PEN_BUTTON_UP
	EVENT_PEN_AXIS                  EventType = C.SDL_EVENT_PEN_AXIS
	EVENT_DROP_BEGIN                EventType = C.SDL_EVENT_DROP_BEGIN
	EVENT_DROP_FILE                 EventType = C.SDL_EVENT_DROP_FILE
	EVENT_DROP_TEXT                 EventType = C.SDL_EVENT_DROP_TEXT
	EVENT_DROP_COMPLETE             EventType = C.SDL_EVENT_DROP_COMPLETE
	EVENT_DROP_POSITION             EventType = C.SDL_EVENT_DROP_POSITION
)

type Event struct {
	Type EventType

	// Keyboard events
	Keyboard *KeyboardEvent

	// Window events
	Window *WindowEvent

	// Mouse events
	MouseMotion *MouseMotionEvent
	MouseButton *MouseButtonEvent
	MouseWheel  *MouseWheelEvent

	// Pen events
	PenProximity *PenProximityEvent
	PenMotion    *PenMotionEvent
	PenTouch     *PenTouchEvent
	PenButton    *PenButtonEvent
	PenAxis      *PenAxisEvent

	// Drag-and-drop events
	Drop *DropEvent
}

func PumpEvents() {
	C.SDL_PumpEvents()
}

func GetTicksNS() uint64 {
	return uint64(C.SDL_GetTicksNS())
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
	case EVENT_WINDOW_RESIZED, EVENT_WINDOW_PIXEL_SIZE_CHANGED:
		event.Window = parseWindowEvent(&cevent)
	case EVENT_KEY_DOWN, EVENT_KEY_UP:
		event.Keyboard = parseKeyboardEvent(&cevent)
	case EVENT_MOUSE_MOTION:
		event.MouseMotion = parseMouseMotionEvent(&cevent)
	case EVENT_MOUSE_BUTTON_DOWN, EVENT_MOUSE_BUTTON_UP:
		event.MouseButton = parseMouseButtonEvent(&cevent)
	case EVENT_MOUSE_WHEEL:
		event.MouseWheel = parseMouseWheelEvent(&cevent)
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

type WindowEvent struct {
	Timestamp uint64
	WindowID  WindowID
	Data1     int32
	Data2     int32
}

func parseWindowEvent(event *C.SDL_Event) *WindowEvent {
	return &WindowEvent{
		Timestamp: uint64(C.get_window_timestamp(event)),
		WindowID:  WindowID(C.get_window_window_id(event)),
		Data1:     int32(C.get_window_data1(event)),
		Data2:     int32(C.get_window_data2(event)),
	}
}
