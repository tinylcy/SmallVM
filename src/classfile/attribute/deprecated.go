package attribute

import "SmallVM/classfile/reader"

type Deprecated struct {
	attributeNameIndex uint16
	attributeLength    uint32
}

func NewDeprecated() *Deprecated {
	return &Deprecated{}
}

func (self *Deprecated) readInfo(reader *reader.ClassReader) {

}
