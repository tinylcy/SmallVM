package constantpool

import "SmallVM/classfile/reader"

type ConstantClassInfo struct {
	index uint16
}

func NewConstantClassInfo() *ConstantClassInfo {
	return &ConstantClassInfo{}
}

func (self *ConstantClassInfo) readInfo(reader *reader.ClassReader) {
	self.index = reader.ReadUInt16()
}
