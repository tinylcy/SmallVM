package reader

import "encoding/binary"

type ClassReader struct {
	data []byte
}

func NewClassReader(data []byte) *ClassReader {
	return &ClassReader{data}
}

func (self *ClassReader) ReadUInt8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

func (self *ClassReader) ReadUInt16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

func (self *ClassReader) ReadUInt32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

func (self *ClassReader) ReadUInt64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

func (self *ClassReader) ReadBytes(n uint32) []byte {
	val := self.data[:n]
	self.data = self.data[n:]
	return val
}
