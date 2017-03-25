package constantpool

import "SmallVM/classfile/reader"

type ConstantMethodRefInfo struct {
	classIndex       uint16
	nameAndTypeIndex uint16
}

func NewConstantMethodRefInfo() *ConstantMethodRefInfo {
	return &ConstantMethodRefInfo{}
}

func (self *ConstantMethodRefInfo) readInfo(reader *reader.ClassReader) {
	self.classIndex = reader.ReadUInt16()
	self.nameAndTypeIndex = reader.ReadUInt16()
}
