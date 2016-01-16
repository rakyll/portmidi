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

// Package portmidi provides PortMidi bindings.
package portmidi

// #cgo LDFLAGS: -lportmidi
// #include <stdlib.h>
// #include <portmidi.h>
// #include <porttime.h>
import "C"

import (
	"errors"
)

// DeviceId is a MIDI device ID.
type DeviceId int

// DeviceInfo provides info about a MIDI device.
type DeviceInfo struct {
	Interface         string
	Name              string
	IsInputAvailable  bool
	IsOutputAvailable bool
	IsOpened          bool
}

type Timestamp int64

// Initializes the portmidi.
func Initialize() error {
	if code := C.Pm_Initialize(); code != 0 {
		return convertToError(code)
	}
	C.Pt_Start(C.int(1), nil, nil)
	return nil
}

// Terminates and cleans up the midi streams.
func Terminate() error {
	C.Pt_Stop()
	return convertToError(C.Pm_Terminate())
}

// Returns the default input device's ID.
func GetDefaultInputDeviceId() DeviceId {
	return DeviceId(C.Pm_GetDefaultInputDeviceID())
}

// Returns the default output device's ID.
func GetDefaultOutputDeviceId() DeviceId {
	return DeviceId(C.Pm_GetDefaultOutputDeviceID())
}

// Returns the number of MIDI devices.
func CountDevices() int {
	return int(C.Pm_CountDevices())
}

// Returns the device info for the device indentified with deviceId.
func GetDeviceInfo(deviceId DeviceId) *DeviceInfo {
	info := C.Pm_GetDeviceInfo(C.PmDeviceID(deviceId))
	return &DeviceInfo{
		Interface:         C.GoString(info.interf),
		Name:              C.GoString(info.name),
		IsInputAvailable:  info.input > 0,
		IsOutputAvailable: info.output > 0,
		IsOpened:          info.opened > 0,
	}
}

// Returns the portmidi timer's current time.
func Time() Timestamp {
	return Timestamp(C.Pt_Time())
}

// convertToError converts a portmidi error code to a Go error.
func convertToError(code C.PmError) error {
	return errors.New(C.GoString(C.Pm_GetErrorText(code)))
}
