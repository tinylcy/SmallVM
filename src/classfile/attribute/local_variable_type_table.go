package attribute

import "SmallVM/classfile/reader"

type LocalVariableTypeTable struct {
	attributeNameIndex           uint16
	attributeLength              uint32
	localVariableTypeTableLength uint16
	localVariableTypeTable       []LocalVariableTypeTableInfo
}

type LocalVariableTypeTableInfo struct {
	startPc        uint16
	length         uint16
	nameIndex      uint16
	signatureIndex uint16
	index          uint16
}

func NewLocalVariableTypeTable(nameIndex uint16) *LocalVariableTypeTable {
	return &LocalVariableTypeTable{attributeNameIndex: nameIndex}
}

func (self *LocalVariableTypeTable) readInfo(reader *reader.ClassReader) {
	self.attributeLength = reader.ReadUInt32()
	self.localVariableTypeTableLength = reader.ReadUInt16()
	self.localVariableTypeTable = self.readLocalVariableTypeTable(reader, self.localVariableTypeTableLength)
}

func (self *LocalVariableTypeTable) readLocalVariableTypeTable(reader *reader.ClassReader, length uint16) []LocalVariableTypeTableInfo {
	var index uint16
	table := make([]LocalVariableTypeTableInfo, length)
	for index = 0; index < length; index++ {
		info := LocalVariableTypeTableInfo{}
		info.startPc = reader.ReadUInt16()
		info.length = reader.ReadUInt16()
		info.nameIndex = reader.ReadUInt16()
		info.signatureIndex = reader.ReadUInt16()
		info.index = reader.ReadUInt16()
		table = append(table, info)
	}
	return table
}
