package attribute

import "SmallVM/classfile/reader"

type Synthetic struct {
	attributeNameIndex uint16
	attributeLength    uint32
}

func NewSynthetic() *Synthetic {
	return &Synthetic{}
}

func (self *Synthetic) readInfo(reader *reader.ClassReader) {

}
