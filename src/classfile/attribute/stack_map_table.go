package attribute

import "SmallVM/classfile/reader"

type StackMapTable struct {
	attributeNameIndex   uint16
	attributeLength      uint32
	numberOfEntries      uint16
	atackMapFrameEntries []StackMapFrame
}

type StackMapFrame struct {
}

func NewStackMapTable() *StackMapTable {
	return &StackMapTable{}
}

func (self *StackMapTable) readInfo(reader *reader.ClassReader) {

}
