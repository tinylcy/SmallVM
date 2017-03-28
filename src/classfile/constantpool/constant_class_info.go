package constantpool

import (
	"SmallVM/classfile/reader"
	"strconv"
)

type ConstantClassInfo struct {
	index uint16
}

func NewConstantClassInfo() *ConstantClassInfo {
	return &ConstantClassInfo{}
}

func (self *ConstantClassInfo) readInfo(reader *reader.ClassReader) {
	self.index = reader.ReadUInt16()
}

func (self *ConstantClassInfo) String() string {
	return strconv.Itoa(int(self.index))
}
