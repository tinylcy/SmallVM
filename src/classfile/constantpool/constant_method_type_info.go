package constantpool

import "SmallVM/classfile/reader"

type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func NewConstantMethodTypeInfo() *ConstantMethodTypeInfo {
	return &ConstantMethodTypeInfo{}
}

func (self *ConstantMethodTypeInfo) readInfo(reader *reader.ClassReader) {
	self.descriptorIndex = reader.ReadUInt16()
}

func (self *ConstantMethodTypeInfo) String() string {
	return ""
}
