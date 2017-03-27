package attribute

import "SmallVM/classfile/reader"

type Signature struct {
	attributeNameIndex uint16
	attributeLength    uint32
	signatureIndex     uint16
}

func NewSignature(nameIndex uint16) *Signature {
	return &Signature{attributeNameIndex: nameIndex}
}

func (self *Signature) readInfo(reader *reader.ClassReader) {
	self.attributeLength = reader.ReadUInt32()
	self.signatureIndex = reader.ReadUInt16()
}
