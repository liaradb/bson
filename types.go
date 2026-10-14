package bson

type (
	// Byte         byte     // 8-bits
	// SignedByte   int8     // 8-bit, two's complement signed integer for which the value is n.
	// UnsignedByte byte     // 8-bit unsigned integer for which the value is n.
	// Int32        int32    // 32-bit signed integer, two's complement
	// Int64        int64    // 64-bit signed integer, two's complement
	// UInt64       uint64   // 64-bit unsigned integer
	// Double       float64  // 64-bit IEEE 754-2008 binary floating point
	Decimel128 [16]byte // 128-bit IEEE 754-2008 decimal floating point
)

type SubType byte

const (
	SubTypeDefault    SubType = 0 //Generic binary subtype. This is the most commonly used binary subtype and should be the 'default' for drivers and tools.
	SubTypeFunction   SubType = 1 //
	SubTypeBinary     SubType = 2 // (old). Deprecated in favor of subtype 0. Drivers and tools should properly handle this subtype. See below for more information.
	SubTypeUUIDLegacy SubType = 3 // (old). Deprecated in favor of subtype 4. Drivers and tools for languages with a native UUID type should properly handle subtype 3.
	SubTypeUUID       SubType = 4 //
	SubTypeMD5        SubType = 5 //
	SubTypeEncrypted  SubType = 6 // BSON value
	SubTypeCompressed SubType = 7 // BSON column. See below for more information.
	SubTypeSensitive  SubType = 8 //
	SubTypeVector     SubType = 9 //. See below for more information.
	// 128 — 255	  User-defined subtypes
)
