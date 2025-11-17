// mouse.go
package sdl3go

/*
#include <SDL3/SDL.h>
#include <SDL3/SDL_mouse.h>

static inline Uint32 get_event_type(SDL_Event *event) {
    return event->type;
}

// Mouse motion event helpers
static inline Uint64 get_motion_timestamp(SDL_Event *event) {
    return event->motion.timestamp;
}

static inline SDL_WindowID get_motion_window(SDL_Event *event) {
    return event->motion.windowID;
}

static inline SDL_MouseID get_motion_which(SDL_Event *event) {
    return event->motion.which;
}

static inline SDL_MouseButtonFlags get_motion_state(SDL_Event *event) {
    return event->motion.state;
}

static inline float get_motion_x(SDL_Event *event) {
    return event->motion.x;
}

static inline float get_motion_y(SDL_Event *event) {
    return event->motion.y;
}

static inline float get_motion_xrel(SDL_Event *event) {
    return event->motion.xrel;
}

static inline float get_motion_yrel(SDL_Event *event) {
    return event->motion.yrel;
}

// Mouse button event helpers
static inline Uint64 get_button_timestamp(SDL_Event *event) {
    return event->button.timestamp;
}

static inline SDL_WindowID get_button_window(SDL_Event *event) {
    return event->button.windowID;
}

static inline SDL_MouseID get_button_which(SDL_Event *event) {
    return event->button.which;
}

static inline Uint8 get_button_button(SDL_Event *event) {
    return event->button.button;
}

static inline bool get_button_down(SDL_Event *event) {
    return event->button.down;
}

static inline Uint8 get_button_clicks(SDL_Event *event) {
    return event->button.clicks;
}

static inline float get_button_x(SDL_Event *event) {
    return event->button.x;
}

static inline float get_button_y(SDL_Event *event) {
    return event->button.y;
}

// Mouse wheel event helpers
static inline Uint64 get_wheel_timestamp(SDL_Event *event) {
    return event->wheel.timestamp;
}

static inline SDL_WindowID get_wheel_window(SDL_Event *event) {
    return event->wheel.windowID;
}

static inline SDL_MouseID get_wheel_which(SDL_Event *event) {
    return event->wheel.which;
}

static inline float get_wheel_x(SDL_Event *event) {
    return event->wheel.x;
}

static inline float get_wheel_y(SDL_Event *event) {
    return event->wheel.y;
}

static inline SDL_MouseWheelDirection get_wheel_direction(SDL_Event *event) {
    return event->wheel.direction;
}

static inline float get_wheel_mouse_x(SDL_Event *event) {
    return event->wheel.mouse_x;
}

static inline float get_wheel_mouse_y(SDL_Event *event) {
    return event->wheel.mouse_y;
}

static inline Sint32 get_wheel_integer_x(SDL_Event *event) {
    return event->wheel.integer_x;
}

static inline Sint32 get_wheel_integer_y(SDL_Event *event) {
    return event->wheel.integer_y;
}
*/
import "C"

// MouseID represents a unique mouse device
type MouseID uint32

// Mouse button constants
const (
	BUTTON_LEFT   = 1
	BUTTON_MIDDLE = 2
	BUTTON_RIGHT  = 3
	BUTTON_X1     = 4
	BUTTON_X2     = 5
)

// Mouse button flags (bitmask for button state)
type MouseButtonFlags uint32

const (
	BUTTON_LMASK  MouseButtonFlags = C.SDL_BUTTON_LMASK
	BUTTON_MMASK  MouseButtonFlags = C.SDL_BUTTON_MMASK
	BUTTON_RMASK  MouseButtonFlags = C.SDL_BUTTON_RMASK
	BUTTON_X1MASK MouseButtonFlags = C.SDL_BUTTON_X1MASK
	BUTTON_X2MASK MouseButtonFlags = C.SDL_BUTTON_X2MASK
)

// Mouse wheel direction
type MouseWheelDirection int

const (
	MOUSEWHEEL_NORMAL  MouseWheelDirection = C.SDL_MOUSEWHEEL_NORMAL
	MOUSEWHEEL_FLIPPED MouseWheelDirection = C.SDL_MOUSEWHEEL_FLIPPED
)

// MouseMotionEvent - Mouse moved
type MouseMotionEvent struct {
	Type      EventType
	Timestamp uint64            // In nanoseconds
	WindowID  WindowID          // The window with mouse focus, if any
	Which     MouseID           // The mouse instance ID
	State     MouseButtonFlags  // The current button state
	X         float32           // X coordinate, relative to window
	Y         float32           // Y coordinate, relative to window
	XRel      float32           // Relative motion in the X direction
	YRel      float32           // Relative motion in the Y direction
}

// MouseButtonEvent - Mouse button pressed or released
type MouseButtonEvent struct {
	Type      EventType
	Timestamp uint64    // In nanoseconds
	WindowID  WindowID  // The window with mouse focus, if any
	Which     MouseID   // The mouse instance ID
	Button    uint8     // The mouse button index (BUTTON_LEFT, etc.)
	Down      bool      // True if button is pressed, false if released
	Clicks    uint8     // 1 for single-click, 2 for double-click, etc.
	X         float32   // X coordinate, relative to window
	Y         float32   // Y coordinate, relative to window
}

// MouseWheelEvent - Mouse wheel scrolled
type MouseWheelEvent struct {
	Type       EventType
	Timestamp  uint64              // In nanoseconds
	WindowID   WindowID            // The window with mouse focus, if any
	Which      MouseID             // The mouse instance ID
	X          float32             // Scroll amount in X direction (horizontal)
	Y          float32             // Scroll amount in Y direction (vertical)
	Direction  MouseWheelDirection // Scroll direction (normal or flipped)
	MouseX     float32             // X coordinate of mouse cursor
	MouseY     float32             // Y coordinate of mouse cursor
	IntegerX   int32               // Accumulated whole scroll ticks in X
	IntegerY   int32               // Accumulated whole scroll ticks in Y
}

// Helper methods for MouseMotionEvent
func (e *MouseMotionEvent) IsLeftButtonDown() bool {
	return (e.State & BUTTON_LMASK) != 0
}

func (e *MouseMotionEvent) IsMiddleButtonDown() bool {
	return (e.State & BUTTON_MMASK) != 0
}

func (e *MouseMotionEvent) IsRightButtonDown() bool {
	return (e.State & BUTTON_RMASK) != 0
}

func (e *MouseMotionEvent) IsX1ButtonDown() bool {
	return (e.State & BUTTON_X1MASK) != 0
}

func (e *MouseMotionEvent) IsX2ButtonDown() bool {
	return (e.State & BUTTON_X2MASK) != 0
}

// Internal parsers for events.go
func parseMouseMotionEvent(cevent *C.SDL_Event) *MouseMotionEvent {
	return &MouseMotionEvent{
		Type:      EVENT_MOUSE_MOTION,
		Timestamp: uint64(C.get_motion_timestamp(cevent)),
		WindowID:  WindowID(C.get_motion_window(cevent)),
		Which:     MouseID(C.get_motion_which(cevent)),
		State:     MouseButtonFlags(C.get_motion_state(cevent)),
		X:         float32(C.get_motion_x(cevent)),
		Y:         float32(C.get_motion_y(cevent)),
		XRel:      float32(C.get_motion_xrel(cevent)),
		YRel:      float32(C.get_motion_yrel(cevent)),
	}
}

func parseMouseButtonEvent(cevent *C.SDL_Event) *MouseButtonEvent {
	eventType := C.get_event_type(cevent)
	return &MouseButtonEvent{
		Type:      EventType(eventType),
		Timestamp: uint64(C.get_button_timestamp(cevent)),
		WindowID:  WindowID(C.get_button_window(cevent)),
		Which:     MouseID(C.get_button_which(cevent)),
		Button:    uint8(C.get_button_button(cevent)),
		Down:      bool(C.get_button_down(cevent)),
		Clicks:    uint8(C.get_button_clicks(cevent)),
		X:         float32(C.get_button_x(cevent)),
		Y:         float32(C.get_button_y(cevent)),
	}
}

func parseMouseWheelEvent(cevent *C.SDL_Event) *MouseWheelEvent {
	return &MouseWheelEvent{
		Type:      EVENT_MOUSE_WHEEL,
		Timestamp: uint64(C.get_wheel_timestamp(cevent)),
		WindowID:  WindowID(C.get_wheel_window(cevent)),
		Which:     MouseID(C.get_wheel_which(cevent)),
		X:         float32(C.get_wheel_x(cevent)),
		Y:         float32(C.get_wheel_y(cevent)),
		Direction: MouseWheelDirection(C.get_wheel_direction(cevent)),
		MouseX:    float32(C.get_wheel_mouse_x(cevent)),
		MouseY:    float32(C.get_wheel_mouse_y(cevent)),
		IntegerX:  int32(C.get_wheel_integer_x(cevent)),
		IntegerY:  int32(C.get_wheel_integer_y(cevent)),
	}
}
