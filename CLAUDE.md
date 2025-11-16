# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

sdl3go is a Go wrapper for SDL3 (Simple DirectMedia Layer 3), designed to make SDL3 easier to use in Go with idiomatic Go patterns. This is primarily being developed to support VALA (Vulkan Animation Linkage Application), a hybrid animation software project. This is not yet a full-blown conversion project, but a targeted wrapper focusing on needed functionality.

**Module:** `github.com/NOT-REAL-GAMES/sdl3go`

## Build and Development Commands

### Building
```bash
# Build the module
go build

# Build with the example
go build ./examples/test.go
```

### Testing
```bash
# Run tests (when test files exist)
go test ./...

# Run specific test
go test -run TestName
```

### Dependencies
Requires SDL3 library installed on the system. The CGo directives are configured for cross-platform linking:
- Windows: `-lSDL3`
- Linux: `-lSDL3`
- macOS: `-lSDL3`

## Architecture

### Core Design Pattern: CGo Wrappers with Type Safety

This codebase wraps SDL3 C functions using CGo, providing type-safe Go interfaces. The architecture follows these patterns:

1. **C Helper Functions**: Complex C union access (like SDL_Event) requires inline C helper functions to extract data safely
2. **Type-Safe Wrappers**: SDL types are wrapped in Go types (e.g., `InitFlags`, `EventType`, `WindowFlags`)
3. **Error Handling**: SDL errors are converted to Go errors via `GetError()` which wraps `SDL_GetError()`
4. **Resource Management**: Go structs (like `Window`) hold C pointers and provide cleanup methods (e.g., `Destroy()`)

### Module Organization

The wrapper is organized by SDL subsystems, each in its own file:

- **sdl.go**: Core SDL initialization, cleanup, and error handling
- **types.go**: Common type definitions (currently `Window` struct)
- **video.go**: Window creation and management
- **vulkan.go**: Vulkan integration (instance extensions, surface creation)
- **events.go**: Event polling and event type routing
- **pen.go**: Comprehensive pen/stylus/tablet input support with all event types

### Event System Architecture

The event system deserves special attention as it's one of the more complex parts:

**Central Event Dispatching** (events.go):
- `PollEvent()` is the main entry point - returns typed Go events
- Uses C helper `get_event_type()` to safely access the event union
- Routes events to subsystem-specific parsers based on type
- Currently handles: QUIT, KEY_DOWN, KEY_UP, and all pen events

**Pen Event Subsystem** (pen.go):
- Comprehensive C helper functions for extracting pen data from SDL_Event union
- Multiple event types: proximity, motion, touch, button, axis
- Each event type has its own Go struct with appropriate fields
- Parser functions (`parsePenMotionEvent`, etc.) convert C events to Go structs
- Helper methods on event structs (e.g., `IsDown()`, `IsEraser()`)

### CGo Patterns Used

**String Handling**:
```go
cTitle := C.CString(title)
defer C.free(unsafe.Pointer(cTitle))
```

**Array Conversion** (see vulkan.go:23):
```go
names := (*[1 << 30]*C.char)(unsafe.Pointer(extensions))[:count:count]
```

**Pointer Passing**:
- SDL handles are stored in Go structs as C pointers
- Methods on Go structs pass the C pointer back to SDL functions
- `unsafe.Pointer` is used for opaque types (Vulkan instance/surface)

## Key Implementation Notes

### Adding New SDL Functions

When wrapping new SDL3 functions:

1. Add CGo imports at the top with appropriate SDL headers
2. If accessing SDL unions (events, etc.), add C helper functions
3. Create Go type wrappers for SDL constants and flags
4. Implement error checking using `GetError()`
5. Handle memory management (C.CString requires explicit free)

### Event System Extension

To add new event types:

1. Add event type constant to `events.go`
2. Create event struct in appropriate subsystem file
3. Add parser function (see pen.go for examples)
4. Add case to switch statement in `PollEvent()`
5. Add field to `Event` struct for the new event type

### Vulkan Integration

The Vulkan integration is minimal but functional:
- `VulkanGetInstanceExtensions()`: Gets required Vulkan instance extensions
- `Window.VulkanCreateSurface()`: Creates VkSurfaceKHR from SDL window
- Uses `unsafe.Pointer` for Vulkan handles to avoid importing full Vulkan bindings

## Development Focus

This wrapper is being developed incrementally based on VALA's needs:
- Prioritize functionality needed for Vulkan-based animation software
- Pen/tablet input is fully implemented (critical for animation work)
- Window and Vulkan integration is functional
- Event system is extensible for additional event types
- Many SDL subsystems are not yet wrapped

When adding features, maintain the established patterns for type safety, error handling, and resource management.
