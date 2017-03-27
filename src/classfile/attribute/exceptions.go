package attribute

import "SmallVM/classfile/reader"

type Exceptions struct {
	attributeNameIndex  uint16
	attributeLength     uint32
	numberOfExceptions  uint16
	exceptionIndexTable []uint16
}

func NewExceptions() *Exceptions {
	return &Exceptions{}
}

func (self *Exceptions) readInfo(reader *reader.ClassReader) {

}
