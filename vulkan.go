// vulkan.go
package sdl3go

/*
#include <SDL3/SDL.h>
#include <SDL3/SDL_vulkan.h>
*/
import "C"
import "unsafe"

func VulkanGetInstanceExtensions() ([]string, error) {
	var count C.Uint32
	if C.SDL_Vulkan_GetInstanceExtensions(&count) == nil {
		return nil, GetError()
	}

	extensions := C.SDL_Vulkan_GetInstanceExtensions(&count)
	if extensions == nil {
		return nil, GetError()
	}

	// Convert C array to Go slice
	names := (*[1 << 30]*C.char)(unsafe.Pointer(extensions))[:count:count]
	result := make([]string, count)
	for i, name := range names {
		result[i] = C.GoString(name)
	}

	return result, nil
}

func (w *Window) VulkanCreateSurface(instance unsafe.Pointer) (unsafe.Pointer, error) {
	var surface C.VkSurfaceKHR

	if !C.SDL_Vulkan_CreateSurface(w.handle, C.VkInstance(instance), nil, &surface) {
		return nil, GetError()
	}

	return unsafe.Pointer(surface), nil
}
