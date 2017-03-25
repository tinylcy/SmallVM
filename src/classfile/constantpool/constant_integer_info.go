package constantpool

import "SmallVM/classfile/reader"

type ConstantIntegerInfo struct {
	val int32
}

func NewConstantIntegetInfo() *ConstantIntegerInfo {
	return &ConstantIntegerInfo{}
}

func (self *ConstantIntegerInfo) readInfo(reader *reader.ClassReader) {
	bytes := reader.ReadUInt32()
	self.val = int32(bytes)
}
