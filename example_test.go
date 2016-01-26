package portmidi_test

import (
	"fmt"
	"log"

	"github.com/rakyll/portmidi"
)

func ExampleWriteSysEx() {
	out, err := portmidi.NewOutputStream(portmidi.GetDefaultOutputDeviceId(), 1024, 0)
	if err != nil {
		log.Fatal(err)
	}

	if err = out.WriteSysEx(portmidi.Time(), "F0 0A 0A 1B 00 7F 30 F7"); err != nil {
		log.Fatal(err)
	}
}

func ExampleWriteSysExBytes() {
	out, err := portmidi.NewOutputStream(portmidi.GetDefaultOutputDeviceId(), 1024, 0)
	if err != nil {
		log.Fatal(err)
	}

	if err = out.WriteSysExBytes(portmidi.Time(), []byte{0xF0, 0x0A, 0x0A, 0x1B, 0x00, 0x7F, 0x30, 0xF7}); err != nil {
		log.Fatal(err)
	}
}

func ExampleReadSysExBytes() {
	in, err := portmidi.NewInputStream(portmidi.GetDefaultInputDeviceId(), 1024)
	if err != nil {
		log.Fatal(err)
	}

	msg, err := in.Read(1024)
	if err != nil {
		log.Fatal(err)
	}

	for i, b := range msg {
		fmt.Printf("SysEx message byte %d = %02x\n", i, b)
	}
}
