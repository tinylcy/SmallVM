package attribute

import "SmallVM/classfile/reader"

type SourceFile struct {
	attributeNameIndex uint16
	attributeLength    uint32
	sourceFileIndex    uint16
}

func NewSourceFile(nameIndex uint16) *SourceFile {
	return &SourceFile{attributeNameIndex: nameIndex}
}

func (self *SourceFile) readInfo(reader *reader.ClassReader) {
	self.attributeLength = reader.ReadUInt32()
	self.sourceFileIndex = reader.ReadUInt16()
}
