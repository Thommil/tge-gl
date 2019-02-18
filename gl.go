// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

import (
	binary "encoding/binary"
	math "math"
	unsafe "unsafe"
)

// Name name of the plugin
const Name = "gl"

var nativeEndian binary.ByteOrder

func init() {
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)
	switch buf {
	case [2]byte{0xAB, 0xCD}:
		nativeEndian = binary.BigEndian
	default:
		nativeEndian = binary.LittleEndian
	}
}

// Uint16ToBytes convert uint16 array to byte array
func Uint16ToBytes(values []uint16) []byte {
	b := make([]byte, 2*len(values))

	u := make([]byte, 2)
	for i, v := range values {
		nativeEndian.PutUint16(u, v)
		b[2*i] = u[0]
		b[2*i+1] = u[1]
	}

	return b
}

// Uint32ToBytes convert uint32 array to byte array
func Uint32ToBytes(values []uint32) []byte {
	b := make([]byte, 4*len(values))

	u := make([]byte, 4)
	for i, v := range values {
		nativeEndian.PutUint32(u, v)
		b[4*i] = u[0]
		b[4*i+1] = u[1]
		b[4*i+2] = u[2]
		b[4*i+3] = u[3]
	}

	return b
}

// Float32ToBytes convert float32 array to byte array
// depending on host endianness
func Float32ToBytes(values []float32) []byte {
	b := make([]byte, 4*len(values))

	if nativeEndian == binary.LittleEndian {
		for i, v := range values {
			u := math.Float32bits(v)
			b[4*i+0] = byte(u >> 0)
			b[4*i+1] = byte(u >> 8)
			b[4*i+2] = byte(u >> 16)
			b[4*i+3] = byte(u >> 24)
		}
	} else {
		for i, v := range values {
			u := math.Float32bits(v)
			b[4*i+0] = byte(u >> 24)
			b[4*i+1] = byte(u >> 16)
			b[4*i+2] = byte(u >> 8)
			b[4*i+3] = byte(u >> 0)
		}
	}

	return b
}

// Float64ToBytes convert float64 array to byte array
// depending on host endianness
func Float64ToBytes(values []float64) []byte {
	b := make([]byte, 8*len(values))

	if nativeEndian == binary.LittleEndian {
		for i, v := range values {
			u := math.Float64bits(v)
			b[8*i+0] = byte(u >> 0)
			b[8*i+1] = byte(u >> 8)
			b[8*i+2] = byte(u >> 16)
			b[8*i+3] = byte(u >> 24)
			b[8*i+4] = byte(u >> 32)
			b[8*i+5] = byte(u >> 40)
			b[8*i+6] = byte(u >> 48)
			b[8*i+7] = byte(u >> 56)
		}
	} else {
		for i, v := range values {
			u := math.Float64bits(v)
			b[8*i+0] = byte(u >> 56)
			b[8*i+1] = byte(u >> 48)
			b[8*i+2] = byte(u >> 40)
			b[8*i+3] = byte(u >> 32)
			b[8*i+4] = byte(u >> 24)
			b[8*i+5] = byte(u >> 16)
			b[8*i+6] = byte(u >> 8)
			b[8*i+7] = byte(u >> 0)
		}
	}

	return b
}
