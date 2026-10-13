package bson

type (
	Byte         byte     // 8-bits
	SignedByte   int8     // 8-bit, two's complement signed integer for which the value is n.
	UnsignedByte byte     // 8-bit unsigned integer for which the value is n.
	Int32        int32    // 32-bit signed integer, two's complement
	Int64        int64    // 64-bit signed integer, two's complement
	UInt64       uint64   // 64-bit unsigned integer
	Double       float64  // 64-bit IEEE 754-2008 binary floating point
	Decimel128   [16]byte // 128-bit IEEE 754-2008 decimal floating point
)

type Element int8

const (
	ElementNone          Element = 0
	ElementDouble        Element = 1   // 64-bit binary floating point
	ElementString        Element = 2   // UTF-8 string
	ElementDocument      Element = 3   // Embedded document
	ElementArray         Element = 4   // Array. See below for more information.
	ElementBinary        Element = 5   // Binary data
	ElementUndefined     Element = 6   // (value). Deprecated.
	ElementObjectID      Element = 7   // ObjectId
	ElementBool          Element = 8   // Boolean
	ElementDateTime      Element = 9   // UTC datetime. int64 is UTC milliseconds since the Unix epoch.
	ElementNull          Element = 10  // value
	ElementRexExp        Element = 11  // Regular expression. See below for more information.
	ElementDBPointer     Element = 12  // (byte*12)	DBPointer. Deprecated.
	ElementCode          Element = 13  // JavaScript code
	ElementSymbol        Element = 14  // Symbol. Deprecated.
	ElementCodeWithScope Element = 15  // JavaScript code with scope. Deprecated.
	ElementInt32         Element = 16  // 32-bit integer
	ElementTimestamp     Element = 17  // Timestamp. A special internal type used by MongoDB replication and sharding. The first 4 bytes are an increment and the second 4 bytes are a timestamp.
	ElementInt64         Element = 18  // 64-bit integer
	ElementDecimal128    Element = 19  // 128-bit decimal floating point
	ElementMinKey        Element = -1  // A special type that compares lower than all other possible BSON element values.
	ElementMaxKey        Element = 127 // A special type that compares higher than all other possible BSON element values.
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
