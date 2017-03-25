package constantpool

import "SmallVM/classfile/reader"

type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func NewConstantMethodHandleInfo() *ConstantMethodHandleInfo {
	return &ConstantMethodHandleInfo{}
}

func (self *ConstantMethodHandleInfo) readInfo(reader *reader.ClassReader) {
	self.referenceKind = reader.ReadUInt8()
	self.referenceIndex = reader.ReadUInt16()
}
