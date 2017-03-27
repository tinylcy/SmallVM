package attribute

import "SmallVM/classfile/reader"

type Signature struct {
	attributeNameIndex uint16
	attributeLength    uint32
	signatureIndex     uint16
}

func NewSignature() *Signature {
	return &Signature{}
}

func (self *Signature) readInfo(reader *reader.ClassReader) {

}
