package constantpool

import "SmallVM/classfile/reader"

type ConstantLongInfo struct {
	val int64
}

func NewConstantLongInfo() *ConstantLongInfo {
	return &ConstantLongInfo{}
}

func (self *ConstantLongInfo) readInfo(reader *reader.ClassReader) {
	bytes := reader.ReadUInt64()
	self.val = int64(bytes)
}
