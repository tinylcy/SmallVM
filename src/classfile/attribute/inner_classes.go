package attribute

import "SmallVM/classfile/reader"

type InnerClasses struct {
	attributeNameIndex uint16
	attributeLength    uint32
	numberOfClasses    uint16
	innerClasses       []InnerClassesInfo
}

type InnerClassesInfo struct {
	innerClassInfoIndex   uint16
	outerClassInfoIndex   uint16
	innerNameIndex        uint16
	innerClassAccessFlags uint16
}

func NewInnerClasses() *InnerClasses {
	return &InnerClasses{}
}

func (self *InnerClasses) readInfo(reader *reader.ClassReader) {

}
