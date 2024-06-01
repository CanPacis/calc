package main

import (
	"bytes"
	"encoding/binary"
	"math"
)

func uint32_to_bytes(n uint32) []byte {
	result := make([]byte, 4)
	binary.LittleEndian.PutUint32(result, n)
	return result
}

func bytes_to_uint32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}

func float64_to_bytes(n float64) []byte {
	buffer := bytes.Buffer{}
	binary.Write(&buffer, binary.LittleEndian, n)
	return buffer.Bytes()
}

func bytes_to_float64(b []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(b))
}
