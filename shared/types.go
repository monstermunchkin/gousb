package shared

// Copyright 2013 Google Inc.  All rights reserved.
// Copyright 2016 the gousb Authors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"fmt"
	"strconv"
)

// ID represents a vendor or product ID.
type ID uint16

// String returns a hexadecimal ID.
func (id ID) String() string {
	return fmt.Sprintf("%04x", int(id))
}

// Class represents a USB-IF (Implementers Forum) class or subclass code.
type Class uint8

// Standard classes defined by USB spec, see https://www.usb.org/defined-class-codes
const (
	ClassPerInterface       Class = 0x00
	ClassAudio              Class = 0x01
	ClassComm               Class = 0x02
	ClassHID                Class = 0x03
	ClassPhysical           Class = 0x05
	ClassImage              Class = 0x06
	ClassPTP                Class = ClassImage // legacy name for image
	ClassPrinter            Class = 0x07
	ClassMassStorage        Class = 0x08
	ClassHub                Class = 0x09
	ClassData               Class = 0x0a
	ClassSmartCard          Class = 0x0b
	ClassContentSecurity    Class = 0x0d
	ClassVideo              Class = 0x0e
	ClassPersonalHealthcare Class = 0x0f
	ClassAudioVideo         Class = 0x10
	ClassBillboard          Class = 0x11
	ClassUSBTypeCBridge     Class = 0x12
	ClassDiagnosticDevice   Class = 0xdc
	ClassWireless           Class = 0xe0
	ClassMiscellaneous      Class = 0xef
	ClassApplication        Class = 0xfe
	ClassVendorSpec         Class = 0xff
)

var classDescription = map[Class]string{
	ClassPerInterface:       "per-interface",
	ClassAudio:              "audio",
	ClassComm:               "communications",
	ClassHID:                "human interface device",
	ClassPhysical:           "physical",
	ClassImage:              "image",
	ClassPrinter:            "printer",
	ClassMassStorage:        "mass storage",
	ClassHub:                "hub",
	ClassData:               "data",
	ClassSmartCard:          "smart card",
	ClassContentSecurity:    "content security",
	ClassVideo:              "video",
	ClassPersonalHealthcare: "personal healthcare",
	ClassAudioVideo:         "audio/video",
	ClassBillboard:          "billboard",
	ClassUSBTypeCBridge:     "USB type-C bridge",
	ClassDiagnosticDevice:   "diagnostic device",
	ClassWireless:           "wireless",
	ClassMiscellaneous:      "miscellaneous",
	ClassApplication:        "application-specific",
	ClassVendorSpec:         "vendor-specific",
}

func (c Class) String() string {
	if d, ok := classDescription[c]; ok {
		return d
	}
	return strconv.Itoa(int(c))
}

// Protocol is the interface class protocol, qualified by the values
// of interface class and subclass.
type Protocol uint8

func (p Protocol) String() string {
	return strconv.Itoa(int(p))
}
