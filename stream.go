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

// #cgo LDFLAGS: -lportmidi -lporttime
// #include <stdlib.h>
// #include <portmidi.h>
// #include <porttime.h>
import "C"

import (
	"errors"
	"unsafe"
)

const (
	minEventBufferSize = 1
	maxEventBufferSize = 1024
)

var (
	errMaxBuffer = errors.New("portmidi: max event buffer size is 1024")
	errMinBuffer = errors.New("portmidi: min event buffer size is 1")
)

// Stream represents a portmidi stream.
type Stream struct {
	deviceId DeviceId
	pmStream *C.PmStream
}

// Event represents a MIDI event.
type Event struct {
	Timestamp Timestamp
	Status    int64
	Data1     int64
	Data2     int64
}

// Initializes a new input stream.
func NewInputStream(deviceId DeviceId, bufferSize int64) (stream *Stream, err error) {
	var str *C.PmStream
	errCode := C.Pm_OpenInput(
		(*unsafe.Pointer)(unsafe.Pointer(&str)),
		C.PmDeviceID(deviceId), nil, C.int32_t(bufferSize), nil, nil)
	if errCode != 0 {
		return nil, convertToError(errCode)
	}
	return &Stream{deviceId: deviceId, pmStream: str}, nil
}

// Initializes a new output stream.
func NewOutputStream(deviceId DeviceId, bufferSize int64, latency int64) (stream *Stream, err error) {
	var str *C.PmStream
	errCode := C.Pm_OpenOutput(
		(*unsafe.Pointer)(unsafe.Pointer(&str)),
		C.PmDeviceID(deviceId), nil, C.int32_t(bufferSize), nil, nil, C.int32_t(latency))
	if errCode != 0 {
		return nil, convertToError(errCode)
	}
	return &Stream{deviceId: deviceId, pmStream: str}, nil
}

// Closes the PortMidi stream.
func (s *Stream) Close() error {
	if s.pmStream == nil {
		return nil
	}
	return convertToError(C.Pm_Close(unsafe.Pointer(s.pmStream)))
}

// Aborts the PortMidi stream.
func (s *Stream) Abort() error {
	if s.pmStream == nil {
		return nil
	}
	return convertToError(C.Pm_Abort(unsafe.Pointer(s.pmStream)))
}

// Writes to the stream.
func (s *Stream) Write(events []Event) error {
	size := len(events)
	if size > maxEventBufferSize {
		return errMaxBuffer
	}
	var buffer []C.PmEvent = make([]C.PmEvent, size)
	for i, evt := range events {
		var event C.PmEvent
		event.timestamp = C.PmTimestamp(evt.Timestamp)
		event.message = C.PmMessage((((evt.Data2 << 16) & 0xFF0000) | ((evt.Data1 << 8) & 0xFF00) | (evt.Status & 0xFF)))
		buffer[i] = event
	}
	return convertToError(C.Pm_Write(unsafe.Pointer(s.pmStream), &buffer[0], C.int32_t(size)))
}

// Writes a MIDI event of three bytes immediately to the stream.
func (s *Stream) WriteShort(status int64, data1 int64, data2 int64) error {
	evt := Event{
		Timestamp: Timestamp(C.Pt_Time()),
		Status:    status,
		Data1:     data1,
		Data2:     data2,
	}
	return s.Write([]Event{evt})
}

func (s *Stream) WriteSysEx(when Timestamp, msg string) error {
	panic("not implemented")
}

func (s *Stream) SetChannelMask(mask int) error {
	panic("not implemented")
}

func (s *Stream) Read(max int) (events []Event, err error) {
	if max > maxEventBufferSize {
		return nil, errMaxBuffer
	}
	if max < minEventBufferSize {
		return nil, errMinBuffer
	}
	buffer := make([]C.PmEvent, max)
	numEvents := C.Pm_Read(unsafe.Pointer(s.pmStream), &buffer[0], C.int32_t(max))
	events = make([]Event, numEvents)
	for i := 0; i < int(numEvents); i++ {
		events[i] = Event{
			Timestamp: Timestamp(buffer[i].timestamp),
			Status:    (int64(buffer[i].message) >> 8) & 0xFF,
			Data1:     (int64(buffer[i].message) >> 16) & 0xFF,
			Data2:     (int64(buffer[i].message) >> 24) & 0xFF,
		}
	}
	return
}

// TODO: add bindings for Pm_SetFilter
