// keyboard_example.go - Demonstrates SDL3 keyboard input
package main

import (
	"fmt"
	"github.com/NOT-REAL-GAMES/sdl3go"
)

func main() {
	// Initialize SDL
	if err := sdl3go.Init(sdl3go.INIT_VIDEO); err != nil {
		panic(fmt.Sprintf("Failed to initialize SDL: %v", err))
	}
	defer sdl3go.Quit()

	// Create a window
	window, err := sdl3go.CreateWindow(
		"Keyboard Input Example",
		800, 600,
		sdl3go.WINDOW_RESIZABLE,
	)
	if err != nil {
		panic(fmt.Sprintf("Failed to create window: %v", err))
	}
	defer window.Destroy()

	fmt.Println("Keyboard Input Example")
	fmt.Println("======================")
	fmt.Println("Try pressing keys with modifiers:")
	fmt.Println("  - Arrow keys to see scancode detection")
	fmt.Println("  - Letter keys with Shift/Ctrl/Alt")
	fmt.Println("  - Press Ctrl+Q or Escape to quit")
	fmt.Println()

	// Main event loop
	running := true
	for running {
		for event, ok := sdl3go.PollEvent(); ok; event, ok = sdl3go.PollEvent() {
			switch event.Type {
			case sdl3go.EVENT_QUIT:
				running = false

			case sdl3go.EVENT_KEY_DOWN:
				if event.Keyboard != nil {
					handleKeyDown(event.Keyboard)

					// Quit on Ctrl+Q or Escape
					if event.Keyboard.Scancode == sdl3go.SCANCODE_Q && event.Keyboard.HasCtrl() {
						fmt.Println("Ctrl+Q pressed - quitting...")
						running = false
					}
					if event.Keyboard.Scancode == sdl3go.SCANCODE_ESCAPE {
						fmt.Println("Escape pressed - quitting...")
						running = false
					}
				}

			case sdl3go.EVENT_KEY_UP:
				if event.Keyboard != nil {
					handleKeyUp(event.Keyboard)
				}
			}
		}
	}
}

func handleKeyDown(key *sdl3go.KeyboardEvent) {
	// Skip key repeats for cleaner output
	if key.Repeat {
		return
	}

	// Build modifier string
	mods := ""
	if key.HasCtrl() {
		mods += "Ctrl+"
	}
	if key.HasShift() {
		mods += "Shift+"
	}
	if key.HasAlt() {
		mods += "Alt+"
	}
	if key.HasGUI() {
		mods += "GUI+"
	}

	// Get key name
	keyName := getKeyName(key.Scancode)

	fmt.Printf("KEY DOWN: %s%s", mods, keyName)

	// Show additional info for interesting keys
	if key.HasCapsLock() {
		fmt.Print(" [CapsLock ON]")
	}
	if key.HasNumLock() {
		fmt.Print(" [NumLock ON]")
	}
	if key.IsModifierOnly() {
		fmt.Print(" (modifier key)")
	}

	fmt.Println()
}

func handleKeyUp(key *sdl3go.KeyboardEvent) {
	// Only show release for modifier keys to reduce noise
	if key.IsModifierOnly() {
		keyName := getKeyName(key.Scancode)
		fmt.Printf("KEY UP: %s (modifier released)\n", keyName)
	}
}

// Helper function to get a readable key name from scancode
func getKeyName(scancode sdl3go.Scancode) string {
	// Letters
	if scancode >= sdl3go.SCANCODE_A && scancode <= sdl3go.SCANCODE_Z {
		return string('A' + rune(scancode-sdl3go.SCANCODE_A))
	}

	// Numbers
	if scancode >= sdl3go.SCANCODE_1 && scancode <= sdl3go.SCANCODE_9 {
		return string('1' + rune(scancode-sdl3go.SCANCODE_1))
	}
	if scancode == sdl3go.SCANCODE_0 {
		return "0"
	}

	// Function keys
	if scancode >= sdl3go.SCANCODE_F1 && scancode <= sdl3go.SCANCODE_F12 {
		return fmt.Sprintf("F%d", int(scancode-sdl3go.SCANCODE_F1)+1)
	}

	// Common keys
	switch scancode {
	case sdl3go.SCANCODE_SPACE:
		return "Space"
	case sdl3go.SCANCODE_RETURN:
		return "Enter"
	case sdl3go.SCANCODE_ESCAPE:
		return "Escape"
	case sdl3go.SCANCODE_BACKSPACE:
		return "Backspace"
	case sdl3go.SCANCODE_TAB:
		return "Tab"
	case sdl3go.SCANCODE_DELETE:
		return "Delete"

	// Arrow keys
	case sdl3go.SCANCODE_UP:
		return "Up"
	case sdl3go.SCANCODE_DOWN:
		return "Down"
	case sdl3go.SCANCODE_LEFT:
		return "Left"
	case sdl3go.SCANCODE_RIGHT:
		return "Right"

	// Modifiers
	case sdl3go.SCANCODE_LSHIFT:
		return "Left Shift"
	case sdl3go.SCANCODE_RSHIFT:
		return "Right Shift"
	case sdl3go.SCANCODE_LCTRL:
		return "Left Ctrl"
	case sdl3go.SCANCODE_RCTRL:
		return "Right Ctrl"
	case sdl3go.SCANCODE_LALT:
		return "Left Alt"
	case sdl3go.SCANCODE_RALT:
		return "Right Alt"
	case sdl3go.SCANCODE_LGUI:
		return "Left GUI"
	case sdl3go.SCANCODE_RGUI:
		return "Right GUI"

	// Special keys
	case sdl3go.SCANCODE_CAPSLOCK:
		return "CapsLock"
	case sdl3go.SCANCODE_NUMLOCKCLEAR:
		return "NumLock"
	case sdl3go.SCANCODE_SCROLLLOCK:
		return "ScrollLock"

	default:
		return fmt.Sprintf("Scancode(%d)", scancode)
	}
}
