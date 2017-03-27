package attribute

import "SmallVM/classfile/reader"

type LocalVariableTable struct {
	attributeNameIndex       uint16
	attributeLength          uint32
	localVariableTableLength uint16
	localVariableTable       []LocalVariableInfo
}

type LocalVariableInfo struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

func NewLocalVariableTable(nameIndex uint16) *LocalVariableTable {
	return &LocalVariableTable{attributeNameIndex: nameIndex}
}

func (self *LocalVariableTable) readInfo(reader *reader.ClassReader) {
	self.attributeLength = reader.ReadUInt32()
	self.localVariableTableLength = reader.ReadUInt16()
	self.localVariableTable = self.readLocalVariableTable(reader, self.localVariableTableLength)
}

func (self *LocalVariableTable) readLocalVariableTable(reader *reader.ClassReader, length uint16) []LocalVariableInfo {
	table := make([]LocalVariableInfo, 0)
	for index := 0; index < int(length); index++ {
		info := LocalVariableInfo{}
		info.startPc = reader.ReadUInt16()
		info.length = reader.ReadUInt16()
		info.nameIndex = reader.ReadUInt16()
		info.descriptorIndex = reader.ReadUInt16()
		info.index = reader.ReadUInt16()
		table = append(table, info)
	}
	return table
}
