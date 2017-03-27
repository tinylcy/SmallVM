package attribute

import "SmallVM/classfile/reader"

type Deprecated struct {
	attributeNameIndex uint16
	attributeLength    uint32
}

func NewDeprecated(nameIndex uint16) *Deprecated {
	return &Deprecated{attributeNameIndex: nameIndex}
}

func (self *Deprecated) readInfo(reader *reader.ClassReader) {
	self.attributeLength = reader.ReadUInt32()
}
