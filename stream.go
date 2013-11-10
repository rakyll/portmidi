// Copyright 2013 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package portmidi provides portmidi bindings.
package portmidi

// #cgo LDFLAGS: -lportmidi
// #include <stdlib.h>
// #include <portmidi.h>
import "C"

import ()

// Stream represents a portmidi stream.
type Stream struct {
	deviceId DeviceId
	midi     interface{}
	// TODO: a pointer to the stream ID
}

// Initializes a new input stream.
func NewInputStream(deviceId DeviceId, bufferSize int64) (stream *Stream, err error) {
	panic("not implemented")
}

// Initialized a new output stream.
func NewOutputStream(deviceId DeviceId, bufferSize int64, latency int64) (stream *Stream, err error) {
	panic("not implemented")
}

// Closes the PortMidi stream.
func (s *Stream) Close() error {
	if code := Pm_Close(s.midi); code != 0 {
		return convertToError(code)
	}
	return nil
}

// Aborts the PortMidi stream.
func (s *Stream) Abort() error {
	if code != Pm_Abort(s.midi); code != 0 {
		return convertToError(code)
	}
	return nil
}

// Writes the stream.
func (s *Stream) Write(data []byte) error {
	panic("not implemented")
}

// Writes
func (s *Stream) WriteShort(status byte, data1 byte, data2 byte) error {
	panic("not implemented")
}

func (s *Stream) WriteSysEx(when Timestamp, msg string) error {
	panic("not implemented")
}

func (s *Stream) SetChannelMask(mask int) error {
	panic("not implemented")
}

func (s *Stream) Poll() error {
	panic("not implemented")
}

// TODO: add bindings for Pm_Read and Pm_SetFilter
