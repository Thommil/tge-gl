// Copyright (c) 2019 Thomas MILLET. All rights reserved.
package gl

import (
	"testing"
)

func TestPointerToBytes(t *testing.T) {

}

const NB_POLYGONS = 1000

func BenchmarkUint16ToBytes(b *testing.B) {
	uint16Array := make([]uint16, NB_POLYGONS*6)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Uint16ToBytes(uint16Array)
	}
}

func BenchmarkUint32ToBytes(b *testing.B) {
	uint32Array := make([]uint32, NB_POLYGONS*6)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Uint32ToBytes(uint32Array)
	}
}

func BenchmarkFloat32ToBytes(b *testing.B) {
	float32Array := make([]float32, NB_POLYGONS*3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Float32ToBytes(float32Array)
	}
}

func BenchmarkFloat64ToBytes(b *testing.B) {
	float64Array := make([]float64, NB_POLYGONS*3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Float64ToBytes(float64Array)
	}
}

func BenchmarkPointerToBytesByte(b *testing.B) {
	uintByteArray := make([]byte, NB_POLYGONS*6)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PointerToBytes(&uintByteArray[0], len(uintByteArray))
	}
}

func BenchmarkPointerToBytesUINT16(b *testing.B) {
	uint16Array := make([]uint16, NB_POLYGONS*6)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PointerToBytes(&uint16Array[0], len(uint16Array))
	}
}

func BenchmarkPointerToBytesUINT32(b *testing.B) {
	uint32Array := make([]uint32, NB_POLYGONS*6)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PointerToBytes(&uint32Array[0], len(uint32Array))
	}
}

func BenchmarkPointerToBytesFLOAT32(b *testing.B) {
	float32Array := make([]float32, NB_POLYGONS*3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PointerToBytes(&float32Array[0], len(float32Array))
	}
}

func BenchmarkPointerToBytesFLOAT64(b *testing.B) {
	float64Array := make([]float64, NB_POLYGONS*3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PointerToBytes(&float64Array[0], len(float64Array))
	}
}
