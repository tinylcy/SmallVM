package attribute

import "SmallVM/classfile/reader"

type Exceptions struct {
	attributeNameIndex  uint16
	attributeLength     uint32
	numberOfExceptions  uint16
	exceptionIndexTable []uint16
}

func NewExceptions(nameIndex uint16) *Exceptions {
	return &Exceptions{attributeNameIndex: nameIndex}
}

func (self *Exceptions) readInfo(reader *reader.ClassReader) {
	self.attributeLength = reader.ReadUInt32()
	self.numberOfExceptions = reader.ReadUInt16()
	self.exceptionIndexTable = make([]uint16, 0)
	var index uint16
	for index = 0; index < self.numberOfExceptions; index++ {
		self.exceptionIndexTable = append(self.exceptionIndexTable, reader.ReadUInt16())
	}
}
