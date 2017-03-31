package common

type ByteCodeReader struct {
	code []byte
	pc   int
}

func NewByteCodeReader(code []byte, pc int) *ByteCodeReader {
	reader := &ByteCodeReader{}
	reader.code = code
	reader.pc = pc
	return reader
}

func (self *ByteCodeReader) ReadUInt8() uint8 {
	value := self.code[self.pc]
	self.pc++
	return uint8(value)
}

func (self *ByteCodeReader) ReadInt8() int8 {
	value := int8(self.ReadUInt8())
	return value
}

func (self *ByteCodeReader) ReadUInt16() uint16 {
	high := self.ReadUInt8()
	low := self.ReadUInt8()
	value := uint16(high)<<8 | uint16(low)
	return value
}

func (self *ByteCodeReader) ReadInt16() int16 {
	value := self.ReadUInt16()
	return int16(value)
}

func (self *ByteCodeReader) ReadUInt32() uint32 {
	byte1 := uint32(self.ReadUInt8())
	byte2 := uint32(self.ReadUInt8())
	byte3 := uint32(self.ReadUInt8())
	byte4 := uint32(self.ReadUInt8())
	value := (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
	return value
}

func (self *ByteCodeReader) ReadInt32() int32 {
	value := int32(self.ReadUInt32())
	return value
}
