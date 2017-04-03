package attribute

import (
	"SmallVM/classfile/constantpool"
	"SmallVM/classfile/reader"
)

type Code struct {
	pool                 []constantpool.ConstantInfo
	attributeNameIndex   uint16
	attributeLength      uint32
	maxStack             uint16
	maxLocals            uint16
	codeLength           uint32
	code                 []byte
	exceptionTableLength uint16
	exceptionTable       []ExceptionInfo
	attributesCount      uint16
	attributes           []AttributeInfo
}

type ExceptionInfo struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func NewCode(cp []constantpool.ConstantInfo, nameIndex uint16) *Code {
	return &Code{pool: cp, attributeNameIndex: nameIndex}
}

func (self *Code) readInfo(reader *reader.ClassReader) {
	self.attributeLength = reader.ReadUInt32()
	self.maxStack = reader.ReadUInt16()
	self.maxLocals = reader.ReadUInt16()
	self.codeLength = reader.ReadUInt32()
	self.code = reader.ReadBytes(self.codeLength)
	self.exceptionTableLength = reader.ReadUInt16()
	self.exceptionTable = readExceptionTable(reader, self.exceptionTableLength)
	self.attributesCount = reader.ReadUInt16()
	self.attributes = ReadAttributes(reader, self.pool, self.attributesCount)
}

func readExceptionTable(reader *reader.ClassReader, length uint16) []ExceptionInfo {
	infoList := make([]ExceptionInfo, 0)
	for index := 0; index < int(length); index++ {
		info := ExceptionInfo{}
		info.startPc = reader.ReadUInt16()
		info.endPc = reader.ReadUInt16()
		info.handlerPc = reader.ReadUInt16()
		info.catchType = reader.ReadUInt16()
		infoList = append(infoList, info)
	}
	return infoList
}

func (self *Code) MaxStack() uint16 {
	return self.maxStack
}

func (self *Code) MaxLocals() uint16 {
	return self.maxLocals
}

func (self *Code) Code() []byte {
	return self.code
}
