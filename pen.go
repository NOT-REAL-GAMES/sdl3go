// pen.go
package sdl3go

/*
#include <SDL3/SDL.h>
#include <SDL3/SDL_pen.h>

static inline Uint32 get_event_type(SDL_Event *event) {
    return event->type;
}

// Helpers to access pen event data from the union
static inline SDL_PenID get_pen_which(SDL_Event *event, Uint32 type) {
    if (type == SDL_EVENT_PEN_PROXIMITY_IN || type == SDL_EVENT_PEN_PROXIMITY_OUT) {
        return event->pproximity.which;
    } else if (type == SDL_EVENT_PEN_MOTION) {
        return event->pmotion.which;
    } else if (type == SDL_EVENT_PEN_DOWN || type == SDL_EVENT_PEN_UP) {
        return event->ptouch.which;
    } else if (type == SDL_EVENT_PEN_BUTTON_DOWN || type == SDL_EVENT_PEN_BUTTON_UP) {
        return event->pbutton.which;
    } else if (type == SDL_EVENT_PEN_AXIS) {
        return event->paxis.which;
    }
    return 0;
}

static inline float get_pen_x(SDL_Event *event, Uint32 type) {
    if (type == SDL_EVENT_PEN_MOTION) return event->pmotion.x;
    if (type == SDL_EVENT_PEN_DOWN || type == SDL_EVENT_PEN_UP) return event->ptouch.x;
    if (type == SDL_EVENT_PEN_BUTTON_DOWN || type == SDL_EVENT_PEN_BUTTON_UP) return event->pbutton.x;
    if (type == SDL_EVENT_PEN_AXIS) return event->paxis.x;
    return 0.0f;
}

static inline float get_pen_y(SDL_Event *event, Uint32 type) {
    if (type == SDL_EVENT_PEN_MOTION) return event->pmotion.y;
    if (type == SDL_EVENT_PEN_DOWN || type == SDL_EVENT_PEN_UP) return event->ptouch.y;
    if (type == SDL_EVENT_PEN_BUTTON_DOWN || type == SDL_EVENT_PEN_BUTTON_UP) return event->pbutton.y;
    if (type == SDL_EVENT_PEN_AXIS) return event->paxis.y;
    return 0.0f;
}

static inline SDL_PenInputFlags get_pen_state(SDL_Event *event, Uint32 type) {
    if (type == SDL_EVENT_PEN_MOTION) return event->pmotion.pen_state;
    if (type == SDL_EVENT_PEN_DOWN || type == SDL_EVENT_PEN_UP) return event->ptouch.pen_state;
    if (type == SDL_EVENT_PEN_BUTTON_DOWN || type == SDL_EVENT_PEN_BUTTON_UP) return event->pbutton.pen_state;
    if (type == SDL_EVENT_PEN_AXIS) return event->paxis.pen_state;
    return 0;
}

static inline bool get_pen_eraser(SDL_Event *event) {
    return event->ptouch.eraser;
}

static inline bool get_pen_down(SDL_Event *event) {
    return event->ptouch.down;
}

static inline Uint8 get_pen_button(SDL_Event *event) {
    return event->pbutton.button;
}

static inline bool get_pen_button_down(SDL_Event *event) {
    return event->pbutton.down;
}

static inline SDL_PenAxis get_pen_axis_type(SDL_Event *event) {
    return event->paxis.axis;
}

static inline float get_pen_axis_value(SDL_Event *event) {
    return event->paxis.value;
}

static inline SDL_WindowID get_pen_window(SDL_Event *event, Uint32 type) {
    if (type == SDL_EVENT_PEN_PROXIMITY_IN || type == SDL_EVENT_PEN_PROXIMITY_OUT) {
        return event->pproximity.windowID;
    } else if (type == SDL_EVENT_PEN_MOTION) {
        return event->pmotion.windowID;
    } else if (type == SDL_EVENT_PEN_DOWN || type == SDL_EVENT_PEN_UP) {
        return event->ptouch.windowID;
    } else if (type == SDL_EVENT_PEN_BUTTON_DOWN || type == SDL_EVENT_PEN_BUTTON_UP) {
        return event->pbutton.windowID;
    } else if (type == SDL_EVENT_PEN_AXIS) {
        return event->paxis.windowID;
    }
    return 0;
}
*/
import "C"

// PenID represents a unique pen/stylus device
type PenID uint32

// Pen input flags (state bitmask)
type PenInputFlags uint32

const (
	PEN_INPUT_DOWN       PenInputFlags = C.SDL_PEN_INPUT_DOWN
	PEN_INPUT_BUTTON_1   PenInputFlags = C.SDL_PEN_INPUT_BUTTON_1
	PEN_INPUT_BUTTON_2   PenInputFlags = C.SDL_PEN_INPUT_BUTTON_2
	PEN_INPUT_BUTTON_3   PenInputFlags = C.SDL_PEN_INPUT_BUTTON_3
	PEN_INPUT_BUTTON_4   PenInputFlags = C.SDL_PEN_INPUT_BUTTON_4
	PEN_INPUT_BUTTON_5   PenInputFlags = C.SDL_PEN_INPUT_BUTTON_5
	PEN_INPUT_ERASER_TIP PenInputFlags = C.SDL_PEN_INPUT_ERASER_TIP
)

// Pen axis indices
type PenAxis int

const (
	PEN_AXIS_PRESSURE            PenAxis = C.SDL_PEN_AXIS_PRESSURE
	PEN_AXIS_XTILT               PenAxis = C.SDL_PEN_AXIS_XTILT
	PEN_AXIS_YTILT               PenAxis = C.SDL_PEN_AXIS_YTILT
	PEN_AXIS_DISTANCE            PenAxis = C.SDL_PEN_AXIS_DISTANCE
	PEN_AXIS_ROTATION            PenAxis = C.SDL_PEN_AXIS_ROTATION
	PEN_AXIS_SLIDER              PenAxis = C.SDL_PEN_AXIS_SLIDER
	PEN_AXIS_TANGENTIAL_PRESSURE PenAxis = C.SDL_PEN_AXIS_TANGENTIAL_PRESSURE
	PEN_AXIS_COUNT               PenAxis = C.SDL_PEN_AXIS_COUNT
)

// Pen event types
/*const (
	EVENT_PEN_PROXIMITY_IN  EventType = C.SDL_EVENT_PEN_PROXIMITY_IN
	EVENT_PEN_PROXIMITY_OUT EventType = C.SDL_EVENT_PEN_PROXIMITY_OUT
	EVENT_PEN_DOWN          EventType = C.SDL_EVENT_PEN_DOWN
	EVENT_PEN_UP            EventType = C.SDL_EVENT_PEN_UP
	EVENT_PEN_MOTION        EventType = C.SDL_EVENT_PEN_MOTION
	EVENT_PEN_BUTTON_DOWN   EventType = C.SDL_EVENT_PEN_BUTTON_DOWN
	EVENT_PEN_BUTTON_UP     EventType = C.SDL_EVENT_PEN_BUTTON_UP
	EVENT_PEN_AXIS          EventType = C.SDL_EVENT_PEN_AXIS
)*/

// WindowID type
type WindowID uint32

// PenProximityEvent - Pen entered/left detection range
type PenProximityEvent struct {
	Type     EventType
	WindowID WindowID
	Which    PenID
}

// PenMotionEvent - Pen moved
type PenMotionEvent struct {
	Type     EventType
	WindowID WindowID
	Which    PenID
	PenState PenInputFlags
	X        float32
	Y        float32
}

// PenTouchEvent - Pen touched/lifted from surface
type PenTouchEvent struct {
	Type     EventType
	WindowID WindowID
	Which    PenID
	PenState PenInputFlags
	X        float32
	Y        float32
	Eraser   bool
	Down     bool
}

// PenButtonEvent - Pen button pressed/released
type PenButtonEvent struct {
	Type     EventType
	WindowID WindowID
	Which    PenID
	PenState PenInputFlags
	X        float32
	Y        float32
	Button   uint8
	Down     bool
}

// PenAxisEvent - Pen axis (pressure, tilt, etc.) changed
type PenAxisEvent struct {
	Type     EventType
	WindowID WindowID
	Which    PenID
	PenState PenInputFlags
	X        float32
	Y        float32
	Axis     PenAxis
	Value    float32
}

// Helper methods
func (e *PenMotionEvent) IsDown() bool {
	return (e.PenState & PEN_INPUT_DOWN) != 0
}

func (e *PenMotionEvent) IsEraser() bool {
	return (e.PenState & PEN_INPUT_ERASER_TIP) != 0
}

func (e *PenTouchEvent) IsEraser() bool {
	return e.Eraser
}

// Internal parsers for events.go
func parsePenProximityEvent(cevent *C.SDL_Event) *PenProximityEvent {
	eventType := C.get_event_type(cevent) // Access directly!
	return &PenProximityEvent{
		Type:     EventType(eventType),
		WindowID: WindowID(C.get_pen_window(cevent, eventType)),
		Which:    PenID(C.get_pen_which(cevent, eventType)),
	}
}

func parsePenMotionEvent(cevent *C.SDL_Event) *PenMotionEvent {
	eventType := C.get_event_type(cevent)
	return &PenMotionEvent{
		Type:     EventType(eventType),
		WindowID: WindowID(C.get_pen_window(cevent, eventType)),
		Which:    PenID(C.get_pen_which(cevent, eventType)),
		PenState: PenInputFlags(C.get_pen_state(cevent, eventType)),
		X:        float32(C.get_pen_x(cevent, eventType)),
		Y:        float32(C.get_pen_y(cevent, eventType)),
	}
}

func parsePenTouchEvent(cevent *C.SDL_Event) *PenTouchEvent {
	eventType := C.get_event_type(cevent)
	return &PenTouchEvent{
		Type:     EventType(eventType),
		WindowID: WindowID(C.get_pen_window(cevent, eventType)),
		Which:    PenID(C.get_pen_which(cevent, eventType)),
		PenState: PenInputFlags(C.get_pen_state(cevent, eventType)),
		X:        float32(C.get_pen_x(cevent, eventType)),
		Y:        float32(C.get_pen_y(cevent, eventType)),
		Eraser:   bool(C.get_pen_eraser(cevent)),
		Down:     bool(C.get_pen_down(cevent)),
	}
}

func parsePenButtonEvent(cevent *C.SDL_Event) *PenButtonEvent {
	eventType := C.get_event_type(cevent)
	return &PenButtonEvent{
		Type:     EventType(eventType),
		WindowID: WindowID(C.get_pen_window(cevent, eventType)),
		Which:    PenID(C.get_pen_which(cevent, eventType)),
		PenState: PenInputFlags(C.get_pen_state(cevent, eventType)),
		X:        float32(C.get_pen_x(cevent, eventType)),
		Y:        float32(C.get_pen_y(cevent, eventType)),
		Button:   uint8(C.get_pen_button(cevent)),
		Down:     bool(C.get_pen_button_down(cevent)),
	}
}

func parsePenAxisEvent(cevent *C.SDL_Event) *PenAxisEvent {
	eventType := C.get_event_type(cevent)
	return &PenAxisEvent{
		Type:     EventType(eventType),
		WindowID: WindowID(C.get_pen_window(cevent, eventType)),
		Which:    PenID(C.get_pen_which(cevent, eventType)),
		PenState: PenInputFlags(C.get_pen_state(cevent, eventType)),
		X:        float32(C.get_pen_x(cevent, eventType)),
		Y:        float32(C.get_pen_y(cevent, eventType)),
		Axis:     PenAxis(C.get_pen_axis_type(cevent)),
		Value:    float32(C.get_pen_axis_value(cevent)),
	}
}
