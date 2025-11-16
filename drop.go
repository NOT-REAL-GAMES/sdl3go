// drop.go
package sdl3go

/*
#include <SDL3/SDL.h>

static inline Uint32 get_event_type(SDL_Event *event) {
    return event->type;
}

// Helpers to access drop event data from the union
static inline SDL_WindowID get_drop_window(SDL_Event *event) {
    return event->drop.windowID;
}

static inline float get_drop_x(SDL_Event *event) {
    return event->drop.x;
}

static inline float get_drop_y(SDL_Event *event) {
    return event->drop.y;
}

static inline const char* get_drop_source(SDL_Event *event) {
    return event->drop.source;
}

static inline const char* get_drop_data(SDL_Event *event) {
    return event->drop.data;
}

static inline Uint64 get_drop_timestamp(SDL_Event *event) {
    return event->drop.timestamp;
}
*/
import "C"

// DropEvent represents a drag-and-drop event
// Used for all drop event types: BEGIN, FILE, TEXT, COMPLETE, POSITION
type DropEvent struct {
	Type      EventType
	Timestamp uint64   // In nanoseconds
	WindowID  WindowID // The window that was dropped on
	X         float32  // X coordinate, relative to window (not set for BEGIN)
	Y         float32  // Y coordinate, relative to window (not set for BEGIN)
	Source    string   // The source app that sent this drop event (may be empty)
	Data      string   // Text for DROP_TEXT, filename for DROP_FILE, empty for others
}

// Internal parser for events.go
func parseDropEvent(cevent *C.SDL_Event) *DropEvent {
	eventType := C.get_event_type(cevent)

	event := &DropEvent{
		Type:      EventType(eventType),
		Timestamp: uint64(C.get_drop_timestamp(cevent)),
		WindowID:  WindowID(C.get_drop_window(cevent)),
		X:         float32(C.get_drop_x(cevent)),
		Y:         float32(C.get_drop_y(cevent)),
	}

	// Convert C strings to Go strings (handle NULL pointers)
	if source := C.get_drop_source(cevent); source != nil {
		event.Source = C.GoString(source)
	}

	if data := C.get_drop_data(cevent); data != nil {
		event.Data = C.GoString(data)
	}

	return event
}
