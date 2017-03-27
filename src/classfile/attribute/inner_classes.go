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

func NewInnerClasses(nameIndex uint16) *InnerClasses {
	return &InnerClasses{attributeNameIndex: nameIndex}
}

func (self *InnerClasses) readInfo(reader *reader.ClassReader) {
	self.attributeLength = reader.ReadUInt32()
	self.numberOfClasses = reader.ReadUInt16()
	self.innerClasses = self.readInnerClasses(reader, self.numberOfClasses)
}

func (self *InnerClasses) readInnerClasses(reader *reader.ClassReader, num uint16) []InnerClassesInfo {
	var index uint16
	classes := make([]InnerClassesInfo, 0)
	for index = 0; index < num; index++ {
		class := InnerClassesInfo{}
		class.innerClassInfoIndex = reader.ReadUInt16()
		class.outerClassInfoIndex = reader.ReadUInt16()
		class.innerNameIndex = reader.ReadUInt16()
		class.innerClassAccessFlags = reader.ReadUInt16()
		classes = append(classes, class)
	}
	return classes
}
