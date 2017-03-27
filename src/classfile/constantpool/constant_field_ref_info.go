package constantpool

import "SmallVM/classfile/reader"

type ConstantFieldRefInfo struct {
	classIndex       uint16
	nameAndTypeIndex uint16
}

func NewConstantFieldRefInfo() *ConstantFieldRefInfo {
	return &ConstantFieldRefInfo{}
}

func (self *ConstantFieldRefInfo) readInfo(reader *reader.ClassReader) {
	self.classIndex = reader.ReadUInt16()
	self.nameAndTypeIndex = reader.ReadUInt16()
}

func (self *ConstantFieldRefInfo) String() string {
	return ""
}
