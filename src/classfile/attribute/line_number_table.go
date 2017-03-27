package attribute

import "SmallVM/classfile/reader"

type LineNumberTable struct {
	attributeNameIndex    uint16
	attributeLength       uint32
	lineNumberTableLength uint16
	lineNumberTable       []LineNumberInfo
}

type LineNumberInfo struct {
	startPc    uint16
	lineNumber uint16
}

func NewLineNumberTable(nameIndex uint16) *LineNumberTable {
	return &LineNumberTable{attributeNameIndex: nameIndex}
}

func (self *LineNumberTable) readInfo(reader *reader.ClassReader) {
	self.attributeLength = reader.ReadUInt32()
	self.lineNumberTableLength = reader.ReadUInt16()
	self.lineNumberTable = self.readLineNumberTable(reader, self.lineNumberTableLength)
}

func (self *LineNumberTable) readLineNumberTable(reader *reader.ClassReader, count uint16) []LineNumberInfo {
	table := make([]LineNumberInfo, 0)
	for index := 0; index < int(count); index++ {
		info := LineNumberInfo{}
		info.startPc = reader.ReadUInt16()
		info.lineNumber = reader.ReadUInt16()
		table = append(table, info)
	}
	return table
}
