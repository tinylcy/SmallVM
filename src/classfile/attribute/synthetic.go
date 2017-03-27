package attribute

import "SmallVM/classfile/reader"

type Synthetic struct {
	attributeNameIndex uint16
	attributeLength    uint32
}

func NewSynthetic(nameIndex uint16) *Synthetic {
	return &Synthetic{attributeNameIndex: nameIndex}
}

func (self *Synthetic) readInfo(reader *reader.ClassReader) {
	self.attributeLength = reader.ReadUInt32()
}
