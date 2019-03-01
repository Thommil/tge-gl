// Copyright (c) 2019 Thomas MILLET. All rights reserved.
package gl

import (
	binary "encoding/binary"
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

// Byte buffer array singleton, allocate 1mB at startup
var byteArrayBuffer = make([]byte, (1024 * 1024))
var byteArrayBufferExtendFactor = 2

func getByteArrayBuffer(size int) []byte {
	if size > len(byteArrayBuffer) {
		for (1024 * 1024 * byteArrayBufferExtendFactor) < size {
			byteArrayBufferExtendFactor++
		}
		byteArrayBuffer = make([]byte, (1024 * 1024 * byteArrayBufferExtendFactor))
	}
	return byteArrayBuffer[:size]
}

// Uint16ToBytes convert uint16 array to byte array
// depending on host endianness
func Uint16ToBytes(values []uint16) []byte {
	b := getByteArrayBuffer(2 * len(values))

	if nativeEndian == binary.LittleEndian {
		for i, v := range values {
			u := *(*uint16)(unsafe.Pointer(&v))
			b[2*i+0] = byte(u)
			b[2*i+1] = byte(u >> 8)
		}
	} else {
		for i, v := range values {
			u := *(*uint16)(unsafe.Pointer(&v))
			b[2*i+0] = byte(u >> 8)
			b[2*i+1] = byte(u)
		}
	}

	return b
}

// Uint32ToBytes convert uint32 array to byte array
// depending on host endianness
func Uint32ToBytes(values []uint32) []byte {
	b := getByteArrayBuffer(4 * len(values))

	if nativeEndian == binary.LittleEndian {
		for i, v := range values {
			u := *(*uint32)(unsafe.Pointer(&v))
			b[2*i+0] = byte(u)
			b[2*i+1] = byte(u >> 8)
			b[4*i+2] = byte(u >> 16)
			b[4*i+3] = byte(u >> 24)
		}
	} else {
		for i, v := range values {
			u := *(*uint32)(unsafe.Pointer(&v))
			b[4*i+0] = byte(u >> 24)
			b[4*i+1] = byte(u >> 16)
			b[4*i+2] = byte(u >> 8)
			b[4*i+3] = byte(u)
		}
	}

	return b
}

// Int16ToBytes convert int16 array to byte array
// depending on host endianness
func Int16ToBytes(values []uint16) []byte {
	b := getByteArrayBuffer(2 * len(values))

	if nativeEndian == binary.LittleEndian {
		for i, v := range values {
			u := *(*int16)(unsafe.Pointer(&v))
			b[2*i+0] = byte(u)
			b[2*i+1] = byte(u >> 8)
		}
	} else {
		for i, v := range values {
			u := *(*int16)(unsafe.Pointer(&v))
			b[2*i+0] = byte(u >> 8)
			b[2*i+1] = byte(u)
		}
	}

	return b
}

// Int32ToBytes convert int32 array to byte array
// depending on host endianness
func Int32ToBytes(values []uint32) []byte {
	b := getByteArrayBuffer(4 * len(values))

	if nativeEndian == binary.LittleEndian {
		for i, v := range values {
			u := *(*int32)(unsafe.Pointer(&v))
			b[2*i+0] = byte(u)
			b[2*i+1] = byte(u >> 8)
			b[4*i+2] = byte(u >> 16)
			b[4*i+3] = byte(u >> 24)
		}
	} else {
		for i, v := range values {
			u := *(*int32)(unsafe.Pointer(&v))
			b[4*i+0] = byte(u >> 24)
			b[4*i+1] = byte(u >> 16)
			b[4*i+2] = byte(u >> 8)
			b[4*i+3] = byte(u)
		}
	}

	return b
}

// Float32ToBytes convert float32 array to byte array
// depending on host endianness
func Float32ToBytes(values []float32) []byte {
	b := getByteArrayBuffer(4 * len(values))

	if nativeEndian == binary.LittleEndian {
		for i, v := range values {
			u := *(*uint32)(unsafe.Pointer(&v))
			b[4*i+0] = byte(u)
			b[4*i+1] = byte(u >> 8)
			b[4*i+2] = byte(u >> 16)
			b[4*i+3] = byte(u >> 24)
		}
	} else {
		for i, v := range values {
			u := *(*uint32)(unsafe.Pointer(&v))
			b[4*i+0] = byte(u >> 24)
			b[4*i+1] = byte(u >> 16)
			b[4*i+2] = byte(u >> 8)
			b[4*i+3] = byte(u)
		}
	}

	return b
}

// Float64ToBytes convert float64 array to byte array
// depending on host endianness
func Float64ToBytes(values []float64) []byte {
	b := getByteArrayBuffer(8 * len(values))

	if nativeEndian == binary.LittleEndian {
		for i, v := range values {
			u := *(*uint64)(unsafe.Pointer(&v))
			b[8*i+0] = byte(u)
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
			u := *(*uint64)(unsafe.Pointer(&v))
			b[8*i+0] = byte(u >> 56)
			b[8*i+1] = byte(u >> 48)
			b[8*i+2] = byte(u >> 40)
			b[8*i+3] = byte(u >> 32)
			b[8*i+4] = byte(u >> 24)
			b[8*i+5] = byte(u >> 16)
			b[8*i+6] = byte(u >> 8)
			b[8*i+7] = byte(u)
		}
	}

	return b
}

// PointerToBytes allows to revover Byte[] from a pointer, usefull for ports (ex: G3N)
func PointerToBytes(data interface{}, size int) []byte {
	switch data.(type) {
	case *uint8:
		b := getByteArrayBuffer(size)
		for i := 0; i < size; i++ {
			b[i] = *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(data.(*uint8))) + uintptr(i)))
		}
		return b
	case *uint16:
		b := getByteArrayBuffer(2 * size)
		if nativeEndian == binary.LittleEndian {
			for i := 0; i < size; i++ {
				v := *(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(data.(*uint16))) + uintptr(i*4)))
				u := *(*uint16)(unsafe.Pointer(&v))
				b[2*i+0] = byte(u)
				b[2*i+1] = byte(u >> 8)
			}
		} else {
			for i := 0; i < size; i++ {
				v := *(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(data.(*uint16))) + uintptr(i*4)))
				u := *(*uint16)(unsafe.Pointer(&v))
				b[2*i+0] = byte(u >> 8)
				b[2*i+1] = byte(u)
			}
		}
		return b
	case *uint32:
		b := getByteArrayBuffer(4 * size)
		if nativeEndian == binary.LittleEndian {
			for i := 0; i < size; i++ {
				v := *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(data.(*uint32))) + uintptr(i*4)))
				u := *(*uint32)(unsafe.Pointer(&v))
				b[4*i+0] = byte(u)
				b[4*i+1] = byte(u >> 8)
				b[4*i+2] = byte(u >> 16)
				b[4*i+3] = byte(u >> 24)
			}
		} else {
			for i := 0; i < size; i++ {
				v := *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(data.(*uint32))) + uintptr(i*4)))
				u := *(*uint32)(unsafe.Pointer(&v))
				b[4*i+0] = byte(u >> 24)
				b[4*i+1] = byte(u >> 16)
				b[4*i+2] = byte(u >> 8)
				b[4*i+3] = byte(u)
			}
		}
		return b
	case *float32:
		b := getByteArrayBuffer(4 * size)
		if nativeEndian == binary.LittleEndian {
			for i := 0; i < size; i++ {
				v := *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(data.(*float32))) + uintptr(i*4)))
				u := *(*uint32)(unsafe.Pointer(&v))
				b[4*i+0] = byte(u)
				b[4*i+1] = byte(u >> 8)
				b[4*i+2] = byte(u >> 16)
				b[4*i+3] = byte(u >> 24)
			}
		} else {
			for i := 0; i < size; i++ {
				v := *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(data.(*float32))) + uintptr(i*4)))
				u := *(*uint32)(unsafe.Pointer(&v))
				b[4*i+0] = byte(u >> 24)
				b[4*i+1] = byte(u >> 16)
				b[4*i+2] = byte(u >> 8)
				b[4*i+3] = byte(u)
			}
		}
		return b
	case *float64:
		b := getByteArrayBuffer(8 * size)
		if nativeEndian == binary.LittleEndian {
			for i := 0; i < size; i++ {
				v := *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(data.(*float64))) + uintptr(i*8)))
				u := *(*uint64)(unsafe.Pointer(&v))
				b[8*i+0] = byte(u)
				b[8*i+1] = byte(u >> 8)
				b[8*i+2] = byte(u >> 16)
				b[8*i+3] = byte(u >> 24)
				b[8*i+4] = byte(u >> 32)
				b[8*i+5] = byte(u >> 40)
				b[8*i+6] = byte(u >> 48)
				b[8*i+7] = byte(u >> 56)
			}
		} else {
			for i := 0; i < size; i++ {
				v := *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(data.(*float64))) + uintptr(i*8)))
				u := *(*uint64)(unsafe.Pointer(&v))
				b[8*i+0] = byte(u >> 56)
				b[8*i+1] = byte(u >> 48)
				b[8*i+2] = byte(u >> 40)
				b[8*i+3] = byte(u >> 32)
				b[8*i+4] = byte(u >> 24)
				b[8*i+5] = byte(u >> 16)
				b[8*i+6] = byte(u >> 8)
				b[8*i+7] = byte(u)
			}
		}
		return b
	}
	return nil
}
