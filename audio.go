package sdl3go

/*
#include <SDL3/SDL.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type AudioDeviceID uint32

const (
	AUDIO_DEVICE_DEFAULT_PLAYBACK  AudioDeviceID = C.SDL_AUDIO_DEVICE_DEFAULT_PLAYBACK
	AUDIO_DEVICE_DEFAULT_RECORDING AudioDeviceID = C.SDL_AUDIO_DEVICE_DEFAULT_RECORDING
)

type AudioStream struct {
	handle *C.SDL_AudioStream
}

type AudioDevice struct {
	ID   AudioDeviceID
	Name string
}

func OpenPlaybackStream(sampleRate, channels int) (*AudioStream, error) {
	return openAudioStream(AUDIO_DEVICE_DEFAULT_PLAYBACK, sampleRate, channels)
}

func OpenRecordingStream(sampleRate, channels int) (*AudioStream, error) {
	return openAudioStream(AUDIO_DEVICE_DEFAULT_RECORDING, sampleRate, channels)
}

func openAudioStream(device AudioDeviceID, sampleRate, channels int) (*AudioStream, error) {
	if sampleRate <= 0 || channels < 1 || channels > 8 {
		return nil, fmt.Errorf("invalid audio format %d Hz, %d channels", sampleRate, channels)
	}
	spec := C.SDL_AudioSpec{format: C.SDL_AUDIO_F32, channels: C.int(channels), freq: C.int(sampleRate)}
	handle := C.SDL_OpenAudioDeviceStream(C.SDL_AudioDeviceID(device), &spec, nil, nil)
	if handle == nil {
		return nil, GetError()
	}
	return &AudioStream{handle: handle}, nil
}

func PlaybackDevices() []AudioDevice {
	return audioDevices(false)
}

func RecordingDevices() []AudioDevice {
	return audioDevices(true)
}

func audioDevices(recording bool) []AudioDevice {
	var count C.int
	var ids *C.SDL_AudioDeviceID
	if recording {
		ids = C.SDL_GetAudioRecordingDevices(&count)
	} else {
		ids = C.SDL_GetAudioPlaybackDevices(&count)
	}
	if ids == nil || count <= 0 {
		return nil
	}
	defer C.SDL_free(unsafe.Pointer(ids))
	raw := unsafe.Slice(ids, int(count))
	devices := make([]AudioDevice, 0, len(raw))
	for _, id := range raw {
		devices = append(devices, AudioDevice{ID: AudioDeviceID(id), Name: C.GoString(C.SDL_GetAudioDeviceName(id))})
	}
	return devices
}

func (s *AudioStream) Resume() error {
	if s == nil || s.handle == nil {
		return fmt.Errorf("audio stream is closed")
	}
	if !C.SDL_ResumeAudioStreamDevice(s.handle) {
		return GetError()
	}
	return nil
}

func (s *AudioStream) Pause() error {
	if s == nil || s.handle == nil {
		return nil
	}
	if !C.SDL_PauseAudioStreamDevice(s.handle) {
		return GetError()
	}
	return nil
}

func (s *AudioStream) Clear() error {
	if s == nil || s.handle == nil {
		return nil
	}
	if !C.SDL_ClearAudioStream(s.handle) {
		return GetError()
	}
	return nil
}

func (s *AudioStream) QueuedBytes() int {
	if s == nil || s.handle == nil {
		return 0
	}
	return int(C.SDL_GetAudioStreamQueued(s.handle))
}

func (s *AudioStream) AvailableBytes() int {
	if s == nil || s.handle == nil {
		return 0
	}
	return int(C.SDL_GetAudioStreamAvailable(s.handle))
}

func (s *AudioStream) PutFloat32(samples []float32) error {
	if s == nil || s.handle == nil {
		return fmt.Errorf("audio stream is closed")
	}
	if len(samples) == 0 {
		return nil
	}
	if !C.SDL_PutAudioStreamData(s.handle, unsafe.Pointer(&samples[0]), C.int(len(samples)*4)) {
		return GetError()
	}
	return nil
}

func (s *AudioStream) GetFloat32(samples []float32) (int, error) {
	if s == nil || s.handle == nil {
		return 0, fmt.Errorf("audio stream is closed")
	}
	if len(samples) == 0 {
		return 0, nil
	}
	got := int(C.SDL_GetAudioStreamData(s.handle, unsafe.Pointer(&samples[0]), C.int(len(samples)*4)))
	if got < 0 {
		return 0, GetError()
	}
	return got / 4, nil
}

func (s *AudioStream) Destroy() {
	if s == nil || s.handle == nil {
		return
	}
	C.SDL_DestroyAudioStream(s.handle)
	s.handle = nil
}
