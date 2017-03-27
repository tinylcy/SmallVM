package constantpool

import "SmallVM/classfile/reader"
import "math"

type ConstantFloatInfo struct {
	val float32
}

func NewConstantFloatInfo() *ConstantFloatInfo {
	return &ConstantFloatInfo{}
}

func (self *ConstantFloatInfo) readInfo(reader *reader.ClassReader) {
	bytes := reader.ReadUInt32()
	self.val = math.Float32frombits(bytes)
}

func (self *ConstantFloatInfo) String() string {
	return ""
}
