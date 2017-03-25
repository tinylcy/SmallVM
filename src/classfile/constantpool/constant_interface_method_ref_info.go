package constantpool

import "SmallVM/classfile/reader"

type ConstantInterfaceMethodRefInfo struct {
	classIndex       uint16
	nameAndTypeIndex uint16
}

func NewConstantInterfaceMethodRefInfo() *ConstantInterfaceMethodRefInfo {
	return &ConstantInterfaceMethodRefInfo{}
}

func (self *ConstantInterfaceMethodRefInfo) readInfo(reader *reader.ClassReader) {
	self.classIndex = reader.ReadUInt16()
	self.nameAndTypeIndex = reader.ReadUInt16()
}
