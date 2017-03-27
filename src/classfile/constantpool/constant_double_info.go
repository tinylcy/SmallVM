package constantpool

import "SmallVM/classfile/reader"
import "math"

type ConstantDoubleInfo struct {
	val float64
}

func NewConstantDoubleInfo() *ConstantDoubleInfo {
	return &ConstantDoubleInfo{}
}

func (self *ConstantDoubleInfo) readInfo(reader *reader.ClassReader) {
	bytes := reader.ReadUInt64()
	self.val = math.Float64frombits(bytes)
}

func (self *ConstantDoubleInfo) String() string {
	return ""
}
