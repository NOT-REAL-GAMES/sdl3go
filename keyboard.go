// keyboard.go
package sdl3go

/*
#include <SDL3/SDL.h>
#include <SDL3/SDL_keyboard.h>

static inline Uint32 get_event_type(SDL_Event *event) {
    return event->type;
}

// Keyboard event helpers
static inline Uint64 get_key_timestamp(SDL_Event *event) {
    return event->key.timestamp;
}

static inline SDL_WindowID get_key_window(SDL_Event *event) {
    return event->key.windowID;
}

static inline SDL_KeyboardID get_key_which(SDL_Event *event) {
    return event->key.which;
}

static inline SDL_Scancode get_key_scancode(SDL_Event *event) {
    return event->key.scancode;
}

static inline SDL_Keycode get_key_key(SDL_Event *event) {
    return event->key.key;
}

static inline SDL_Keymod get_key_mod(SDL_Event *event) {
    return event->key.mod;
}

static inline Uint16 get_key_raw(SDL_Event *event) {
    return event->key.raw;
}

static inline bool get_key_down(SDL_Event *event) {
    return event->key.down;
}

static inline bool get_key_repeat(SDL_Event *event) {
    return event->key.repeat;
}
*/
import "C"

// KeyboardID represents a unique keyboard device
type KeyboardID uint32

// Scancode represents a physical key location on the keyboard
type Scancode uint32

// SDL_Scancode - Physical key scancodes
const (
	SCANCODE_UNKNOWN Scancode = C.SDL_SCANCODE_UNKNOWN

	// Letters
	SCANCODE_A Scancode = C.SDL_SCANCODE_A
	SCANCODE_B Scancode = C.SDL_SCANCODE_B
	SCANCODE_C Scancode = C.SDL_SCANCODE_C
	SCANCODE_D Scancode = C.SDL_SCANCODE_D
	SCANCODE_E Scancode = C.SDL_SCANCODE_E
	SCANCODE_F Scancode = C.SDL_SCANCODE_F
	SCANCODE_G Scancode = C.SDL_SCANCODE_G
	SCANCODE_H Scancode = C.SDL_SCANCODE_H
	SCANCODE_I Scancode = C.SDL_SCANCODE_I
	SCANCODE_J Scancode = C.SDL_SCANCODE_J
	SCANCODE_K Scancode = C.SDL_SCANCODE_K
	SCANCODE_L Scancode = C.SDL_SCANCODE_L
	SCANCODE_M Scancode = C.SDL_SCANCODE_M
	SCANCODE_N Scancode = C.SDL_SCANCODE_N
	SCANCODE_O Scancode = C.SDL_SCANCODE_O
	SCANCODE_P Scancode = C.SDL_SCANCODE_P
	SCANCODE_Q Scancode = C.SDL_SCANCODE_Q
	SCANCODE_R Scancode = C.SDL_SCANCODE_R
	SCANCODE_S Scancode = C.SDL_SCANCODE_S
	SCANCODE_T Scancode = C.SDL_SCANCODE_T
	SCANCODE_U Scancode = C.SDL_SCANCODE_U
	SCANCODE_V Scancode = C.SDL_SCANCODE_V
	SCANCODE_W Scancode = C.SDL_SCANCODE_W
	SCANCODE_X Scancode = C.SDL_SCANCODE_X
	SCANCODE_Y Scancode = C.SDL_SCANCODE_Y
	SCANCODE_Z Scancode = C.SDL_SCANCODE_Z

	// Numbers
	SCANCODE_1 Scancode = C.SDL_SCANCODE_1
	SCANCODE_2 Scancode = C.SDL_SCANCODE_2
	SCANCODE_3 Scancode = C.SDL_SCANCODE_3
	SCANCODE_4 Scancode = C.SDL_SCANCODE_4
	SCANCODE_5 Scancode = C.SDL_SCANCODE_5
	SCANCODE_6 Scancode = C.SDL_SCANCODE_6
	SCANCODE_7 Scancode = C.SDL_SCANCODE_7
	SCANCODE_8 Scancode = C.SDL_SCANCODE_8
	SCANCODE_9 Scancode = C.SDL_SCANCODE_9
	SCANCODE_0 Scancode = C.SDL_SCANCODE_0

	// Function keys
	SCANCODE_RETURN    Scancode = C.SDL_SCANCODE_RETURN
	SCANCODE_ESCAPE    Scancode = C.SDL_SCANCODE_ESCAPE
	SCANCODE_BACKSPACE Scancode = C.SDL_SCANCODE_BACKSPACE
	SCANCODE_TAB       Scancode = C.SDL_SCANCODE_TAB
	SCANCODE_SPACE     Scancode = C.SDL_SCANCODE_SPACE

	// Symbols
	SCANCODE_MINUS        Scancode = C.SDL_SCANCODE_MINUS
	SCANCODE_EQUALS       Scancode = C.SDL_SCANCODE_EQUALS
	SCANCODE_LEFTBRACKET  Scancode = C.SDL_SCANCODE_LEFTBRACKET
	SCANCODE_RIGHTBRACKET Scancode = C.SDL_SCANCODE_RIGHTBRACKET
	SCANCODE_BACKSLASH    Scancode = C.SDL_SCANCODE_BACKSLASH
	SCANCODE_SEMICOLON    Scancode = C.SDL_SCANCODE_SEMICOLON
	SCANCODE_APOSTROPHE   Scancode = C.SDL_SCANCODE_APOSTROPHE
	SCANCODE_GRAVE        Scancode = C.SDL_SCANCODE_GRAVE
	SCANCODE_COMMA        Scancode = C.SDL_SCANCODE_COMMA
	SCANCODE_PERIOD       Scancode = C.SDL_SCANCODE_PERIOD
	SCANCODE_SLASH        Scancode = C.SDL_SCANCODE_SLASH

	SCANCODE_CAPSLOCK Scancode = C.SDL_SCANCODE_CAPSLOCK

	// Function keys F1-F12
	SCANCODE_F1  Scancode = C.SDL_SCANCODE_F1
	SCANCODE_F2  Scancode = C.SDL_SCANCODE_F2
	SCANCODE_F3  Scancode = C.SDL_SCANCODE_F3
	SCANCODE_F4  Scancode = C.SDL_SCANCODE_F4
	SCANCODE_F5  Scancode = C.SDL_SCANCODE_F5
	SCANCODE_F6  Scancode = C.SDL_SCANCODE_F6
	SCANCODE_F7  Scancode = C.SDL_SCANCODE_F7
	SCANCODE_F8  Scancode = C.SDL_SCANCODE_F8
	SCANCODE_F9  Scancode = C.SDL_SCANCODE_F9
	SCANCODE_F10 Scancode = C.SDL_SCANCODE_F10
	SCANCODE_F11 Scancode = C.SDL_SCANCODE_F11
	SCANCODE_F12 Scancode = C.SDL_SCANCODE_F12

	// Special keys
	SCANCODE_PRINTSCREEN Scancode = C.SDL_SCANCODE_PRINTSCREEN
	SCANCODE_SCROLLLOCK  Scancode = C.SDL_SCANCODE_SCROLLLOCK
	SCANCODE_PAUSE       Scancode = C.SDL_SCANCODE_PAUSE
	SCANCODE_INSERT      Scancode = C.SDL_SCANCODE_INSERT
	SCANCODE_HOME        Scancode = C.SDL_SCANCODE_HOME
	SCANCODE_PAGEUP      Scancode = C.SDL_SCANCODE_PAGEUP
	SCANCODE_DELETE      Scancode = C.SDL_SCANCODE_DELETE
	SCANCODE_END         Scancode = C.SDL_SCANCODE_END
	SCANCODE_PAGEDOWN    Scancode = C.SDL_SCANCODE_PAGEDOWN

	// Arrow keys
	SCANCODE_RIGHT Scancode = C.SDL_SCANCODE_RIGHT
	SCANCODE_LEFT  Scancode = C.SDL_SCANCODE_LEFT
	SCANCODE_DOWN  Scancode = C.SDL_SCANCODE_DOWN
	SCANCODE_UP    Scancode = C.SDL_SCANCODE_UP

	// Numpad
	SCANCODE_NUMLOCKCLEAR Scancode = C.SDL_SCANCODE_NUMLOCKCLEAR
	SCANCODE_KP_DIVIDE    Scancode = C.SDL_SCANCODE_KP_DIVIDE
	SCANCODE_KP_MULTIPLY  Scancode = C.SDL_SCANCODE_KP_MULTIPLY
	SCANCODE_KP_MINUS     Scancode = C.SDL_SCANCODE_KP_MINUS
	SCANCODE_KP_PLUS      Scancode = C.SDL_SCANCODE_KP_PLUS
	SCANCODE_KP_ENTER     Scancode = C.SDL_SCANCODE_KP_ENTER
	SCANCODE_KP_1         Scancode = C.SDL_SCANCODE_KP_1
	SCANCODE_KP_2         Scancode = C.SDL_SCANCODE_KP_2
	SCANCODE_KP_3         Scancode = C.SDL_SCANCODE_KP_3
	SCANCODE_KP_4         Scancode = C.SDL_SCANCODE_KP_4
	SCANCODE_KP_5         Scancode = C.SDL_SCANCODE_KP_5
	SCANCODE_KP_6         Scancode = C.SDL_SCANCODE_KP_6
	SCANCODE_KP_7         Scancode = C.SDL_SCANCODE_KP_7
	SCANCODE_KP_8         Scancode = C.SDL_SCANCODE_KP_8
	SCANCODE_KP_9         Scancode = C.SDL_SCANCODE_KP_9
	SCANCODE_KP_0         Scancode = C.SDL_SCANCODE_KP_0
	SCANCODE_KP_PERIOD    Scancode = C.SDL_SCANCODE_KP_PERIOD
	SCANCODE_KP_EQUALS    Scancode = C.SDL_SCANCODE_KP_EQUALS
	SCANCODE_KP_COMMA     Scancode = C.SDL_SCANCODE_KP_COMMA

	// Extended function keys
	SCANCODE_F13 Scancode = C.SDL_SCANCODE_F13
	SCANCODE_F14 Scancode = C.SDL_SCANCODE_F14
	SCANCODE_F15 Scancode = C.SDL_SCANCODE_F15
	SCANCODE_F16 Scancode = C.SDL_SCANCODE_F16
	SCANCODE_F17 Scancode = C.SDL_SCANCODE_F17
	SCANCODE_F18 Scancode = C.SDL_SCANCODE_F18
	SCANCODE_F19 Scancode = C.SDL_SCANCODE_F19
	SCANCODE_F20 Scancode = C.SDL_SCANCODE_F20
	SCANCODE_F21 Scancode = C.SDL_SCANCODE_F21
	SCANCODE_F22 Scancode = C.SDL_SCANCODE_F22
	SCANCODE_F23 Scancode = C.SDL_SCANCODE_F23
	SCANCODE_F24 Scancode = C.SDL_SCANCODE_F24

	// Application and system keys
	SCANCODE_EXECUTE    Scancode = C.SDL_SCANCODE_EXECUTE
	SCANCODE_HELP       Scancode = C.SDL_SCANCODE_HELP
	SCANCODE_MENU       Scancode = C.SDL_SCANCODE_MENU
	SCANCODE_SELECT     Scancode = C.SDL_SCANCODE_SELECT
	SCANCODE_STOP       Scancode = C.SDL_SCANCODE_STOP
	SCANCODE_AGAIN      Scancode = C.SDL_SCANCODE_AGAIN
	SCANCODE_UNDO       Scancode = C.SDL_SCANCODE_UNDO
	SCANCODE_CUT        Scancode = C.SDL_SCANCODE_CUT
	SCANCODE_COPY       Scancode = C.SDL_SCANCODE_COPY
	SCANCODE_PASTE      Scancode = C.SDL_SCANCODE_PASTE
	SCANCODE_FIND       Scancode = C.SDL_SCANCODE_FIND
	SCANCODE_MUTE       Scancode = C.SDL_SCANCODE_MUTE
	SCANCODE_VOLUMEUP   Scancode = C.SDL_SCANCODE_VOLUMEUP
	SCANCODE_VOLUMEDOWN Scancode = C.SDL_SCANCODE_VOLUMEDOWN

	// Modifier keys
	SCANCODE_LCTRL  Scancode = C.SDL_SCANCODE_LCTRL
	SCANCODE_LSHIFT Scancode = C.SDL_SCANCODE_LSHIFT
	SCANCODE_LALT   Scancode = C.SDL_SCANCODE_LALT
	SCANCODE_LGUI   Scancode = C.SDL_SCANCODE_LGUI // Windows/Command/Meta key
	SCANCODE_RCTRL  Scancode = C.SDL_SCANCODE_RCTRL
	SCANCODE_RSHIFT Scancode = C.SDL_SCANCODE_RSHIFT
	SCANCODE_RALT   Scancode = C.SDL_SCANCODE_RALT
	SCANCODE_RGUI   Scancode = C.SDL_SCANCODE_RGUI

	// International and extended keys
	SCANCODE_MODE          Scancode = C.SDL_SCANCODE_MODE
	SCANCODE_SLEEP         Scancode = C.SDL_SCANCODE_SLEEP
	SCANCODE_WAKE          Scancode = C.SDL_SCANCODE_WAKE
	SCANCODE_CHANNEL_INCREMENT Scancode = C.SDL_SCANCODE_CHANNEL_INCREMENT
	SCANCODE_CHANNEL_DECREMENT Scancode = C.SDL_SCANCODE_CHANNEL_DECREMENT

	// Media keys
	SCANCODE_MEDIA_PLAY       Scancode = C.SDL_SCANCODE_MEDIA_PLAY
	SCANCODE_MEDIA_PAUSE      Scancode = C.SDL_SCANCODE_MEDIA_PAUSE
	SCANCODE_MEDIA_RECORD     Scancode = C.SDL_SCANCODE_MEDIA_RECORD
	SCANCODE_MEDIA_FAST_FORWARD Scancode = C.SDL_SCANCODE_MEDIA_FAST_FORWARD
	SCANCODE_MEDIA_REWIND     Scancode = C.SDL_SCANCODE_MEDIA_REWIND
	SCANCODE_MEDIA_NEXT_TRACK Scancode = C.SDL_SCANCODE_MEDIA_NEXT_TRACK
	SCANCODE_MEDIA_PREVIOUS_TRACK Scancode = C.SDL_SCANCODE_MEDIA_PREVIOUS_TRACK
	SCANCODE_MEDIA_STOP       Scancode = C.SDL_SCANCODE_MEDIA_STOP
	SCANCODE_MEDIA_EJECT      Scancode = C.SDL_SCANCODE_MEDIA_EJECT
	SCANCODE_MEDIA_PLAY_PAUSE Scancode = C.SDL_SCANCODE_MEDIA_PLAY_PAUSE
	SCANCODE_MEDIA_SELECT     Scancode = C.SDL_SCANCODE_MEDIA_SELECT

	// Application control
	SCANCODE_AC_NEW        Scancode = C.SDL_SCANCODE_AC_NEW
	SCANCODE_AC_OPEN       Scancode = C.SDL_SCANCODE_AC_OPEN
	SCANCODE_AC_CLOSE      Scancode = C.SDL_SCANCODE_AC_CLOSE
	SCANCODE_AC_EXIT       Scancode = C.SDL_SCANCODE_AC_EXIT
	SCANCODE_AC_SAVE       Scancode = C.SDL_SCANCODE_AC_SAVE
	SCANCODE_AC_PRINT      Scancode = C.SDL_SCANCODE_AC_PRINT
	SCANCODE_AC_PROPERTIES Scancode = C.SDL_SCANCODE_AC_PROPERTIES
	SCANCODE_AC_SEARCH     Scancode = C.SDL_SCANCODE_AC_SEARCH
	SCANCODE_AC_HOME       Scancode = C.SDL_SCANCODE_AC_HOME
	SCANCODE_AC_BACK       Scancode = C.SDL_SCANCODE_AC_BACK
	SCANCODE_AC_FORWARD    Scancode = C.SDL_SCANCODE_AC_FORWARD
	SCANCODE_AC_STOP       Scancode = C.SDL_SCANCODE_AC_STOP
	SCANCODE_AC_REFRESH    Scancode = C.SDL_SCANCODE_AC_REFRESH
	SCANCODE_AC_BOOKMARKS  Scancode = C.SDL_SCANCODE_AC_BOOKMARKS

	// Miscellaneous
	SCANCODE_SOFTLEFT      Scancode = C.SDL_SCANCODE_SOFTLEFT
	SCANCODE_SOFTRIGHT     Scancode = C.SDL_SCANCODE_SOFTRIGHT
	SCANCODE_CALL          Scancode = C.SDL_SCANCODE_CALL
	SCANCODE_ENDCALL       Scancode = C.SDL_SCANCODE_ENDCALL
)

// Keycode represents a virtual key (with layout)
type Keycode uint32

// Keymod represents keyboard modifier flags (Ctrl, Shift, Alt, etc.)
type Keymod uint16

// SDL_Keymod - Keyboard modifier flags
const (
	KMOD_NONE  Keymod = C.SDL_KMOD_NONE  // No modifier
	KMOD_LSHIFT Keymod = C.SDL_KMOD_LSHIFT // Left Shift
	KMOD_RSHIFT Keymod = C.SDL_KMOD_RSHIFT // Right Shift
	KMOD_LCTRL Keymod = C.SDL_KMOD_LCTRL  // Left Ctrl
	KMOD_RCTRL Keymod = C.SDL_KMOD_RCTRL  // Right Ctrl
	KMOD_LALT  Keymod = C.SDL_KMOD_LALT   // Left Alt
	KMOD_RALT  Keymod = C.SDL_KMOD_RALT   // Right Alt
	KMOD_LGUI  Keymod = C.SDL_KMOD_LGUI   // Left GUI (Windows/Command/Meta)
	KMOD_RGUI  Keymod = C.SDL_KMOD_RGUI   // Right GUI
	KMOD_NUM   Keymod = C.SDL_KMOD_NUM    // Num Lock
	KMOD_CAPS  Keymod = C.SDL_KMOD_CAPS   // Caps Lock
	KMOD_MODE  Keymod = C.SDL_KMOD_MODE   // AltGr
	KMOD_SCROLL Keymod = C.SDL_KMOD_SCROLL // Scroll Lock

	// Combined modifiers (convenience)
	KMOD_SHIFT Keymod = C.SDL_KMOD_SHIFT // Any Shift
	KMOD_CTRL  Keymod = C.SDL_KMOD_CTRL  // Any Ctrl
	KMOD_ALT   Keymod = C.SDL_KMOD_ALT   // Any Alt
	KMOD_GUI   Keymod = C.SDL_KMOD_GUI   // Any GUI
)

// KeyboardEvent - Keyboard key pressed or released
type KeyboardEvent struct {
	Type      EventType   // EVENT_KEY_DOWN or EVENT_KEY_UP
	Timestamp uint64      // In nanoseconds
	WindowID  WindowID    // The window with keyboard focus, if any
	Which     KeyboardID  // The keyboard instance ID
	Scancode  Scancode    // Physical key code
	Key       Keycode     // Virtual key code (affected by layout)
	Mod       Keymod      // Active keyboard modifiers
	Raw       uint16      // Platform-specific scancode
	Down      bool        // True if key is pressed, false if released
	Repeat    bool        // True if this is a key repeat
}

// Helper methods for KeyboardEvent

// HasShift returns true if any Shift key is pressed
func (e *KeyboardEvent) HasShift() bool {
	return (e.Mod & KMOD_SHIFT) != 0
}

// HasCtrl returns true if any Ctrl key is pressed
func (e *KeyboardEvent) HasCtrl() bool {
	return (e.Mod & KMOD_CTRL) != 0
}

// HasAlt returns true if any Alt key is pressed
func (e *KeyboardEvent) HasAlt() bool {
	return (e.Mod & KMOD_ALT) != 0
}

// HasGUI returns true if any GUI/Windows/Command key is pressed
func (e *KeyboardEvent) HasGUI() bool {
	return (e.Mod & KMOD_GUI) != 0
}

// HasCapsLock returns true if Caps Lock is active
func (e *KeyboardEvent) HasCapsLock() bool {
	return (e.Mod & KMOD_CAPS) != 0
}

// HasNumLock returns true if Num Lock is active
func (e *KeyboardEvent) HasNumLock() bool {
	return (e.Mod & KMOD_NUM) != 0
}

// IsModifierOnly returns true if this is a modifier key press (Ctrl, Shift, Alt, GUI)
func (e *KeyboardEvent) IsModifierOnly() bool {
	switch e.Scancode {
	case SCANCODE_LCTRL, SCANCODE_RCTRL,
		SCANCODE_LSHIFT, SCANCODE_RSHIFT,
		SCANCODE_LALT, SCANCODE_RALT,
		SCANCODE_LGUI, SCANCODE_RGUI:
		return true
	}
	return false
}

// Internal parser for events.go
func parseKeyboardEvent(cevent *C.SDL_Event) *KeyboardEvent {
	eventType := C.get_event_type(cevent)
	return &KeyboardEvent{
		Type:      EventType(eventType),
		Timestamp: uint64(C.get_key_timestamp(cevent)),
		WindowID:  WindowID(C.get_key_window(cevent)),
		Which:     KeyboardID(C.get_key_which(cevent)),
		Scancode:  Scancode(C.get_key_scancode(cevent)),
		Key:       Keycode(C.get_key_key(cevent)),
		Mod:       Keymod(C.get_key_mod(cevent)),
		Raw:       uint16(C.get_key_raw(cevent)),
		Down:      bool(C.get_key_down(cevent)),
		Repeat:    bool(C.get_key_repeat(cevent)),
	}
}
