# portmidi

This package contains Go bindings for PortMidi.

## Usage
In order to start, go get this repository:
~~~ go
go get github.com/rakyll/portmidi
~~~
`libportmidi` is required as a development dependency.

~~~ go
import (
    "github.com/rakyll/portmidi"
)

portmidi.Initialize()
stream, _ := portmidi.NewOutputStream(deviceId, 1024, 0)

// note on events to play C# minor chord
stream.WriteShort(0x90, 60, 100)
stream.WriteShort(0x90, 64, 100)
stream.WriteShort(0x90, 67, 100)

time.Sleep(2 * time.Second)

// note off events
stream.WriteShort(0x80, 60, 100)
stream.WriteShort(0x80, 64, 100)
stream.WriteShort(0x80, 67, 100)

stream.Close()
portmidi.Terminate()
~~~
    
## License
    Copyright 2013 Google Inc. All Rights Reserved.
    
    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at
    
         http://www.apache.org/licenses/LICENSE-2.0
    
    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
