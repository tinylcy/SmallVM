package attribute

import "SmallVM/classfile/reader"

type Unparsed struct {
	attributeNameIndex uint16
	attributeLength    uint32
	info               []byte
}

func NewUnparsed(nameIndex uint16) *Unparsed {
	return &Unparsed{attributeNameIndex: nameIndex}
}

func (self *Unparsed) readInfo(reader *reader.ClassReader) {
	self.attributeLength = reader.ReadUInt32()
	self.info = reader.ReadBytes(self.attributeLength)
}
