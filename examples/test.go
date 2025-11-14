// examples/test.go
package main

import (
	"fmt"
	"time"

	sdl "github.com/NOT-REAL-GAMES/sdl3go"
)

func main() {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("SDL3 + Vulkan", 1280, 720, sdl.WINDOW_VULKAN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	exts, err := sdl.VulkanGetInstanceExtensions()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Required Vulkan extensions: %v\n", exts)

	time.Sleep(3 * time.Second)

}
