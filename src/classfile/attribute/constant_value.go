package attribute

import "SmallVM/classfile/reader"

type ConstantValue struct {
	attributeNameIndex uint16
	attributeLength    uint32
	constantValueIndex uint16
}

func NewConstantValue(nameIndex uint16) *ConstantValue {
	return &ConstantValue{attributeNameIndex: nameIndex}
}

func (self *ConstantValue) readInfo(reader *reader.ClassReader) {
	self.attributeLength = reader.ReadUInt32()
	self.constantValueIndex = reader.ReadUInt16()
}
