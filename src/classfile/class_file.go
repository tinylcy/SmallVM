package classfile

import "SmallVM/classfile/reader"
import "SmallVM/classfile/constantpool"

type ClassFile struct {
	magic             uint32
	minorVersion      uint16
	majorVersion      uint16
	constantPoolCount uint16
	constantPool      []constantpool.ConstantInfo
	accessFlags       uint16
	thisClass         uint16
	superClass        uint16
	interfacesCount   uint16
	interfaces        []uint16
	fieldsCount       uint16
	fields            []FieldInfo
	methodsCount      uint16
	methods           []MethodInfo
	attributesCount   uint16
	attributes        []AttributeInfo
}

type FieldInfo struct {
}

type MethodInfo struct {
}

type AttributeInfo struct {
}

func NewClassFile() *ClassFile {
	return &ClassFile{}
}

func (self *ClassFile) Parse(reader *reader.ClassReader) (*ClassFile, error) {
	self.ReadMagic(reader)
	self.ReadMinorVersion(reader)
	self.ReadMajorVersion(reader)
	self.ReadConstantPoolCount(reader)
	// parse constant pool info
	self.constantPool = constantpool.Parse(reader, self.constantPoolCount)
	return self, nil
}

func (self *ClassFile) ReadMagic(reader *reader.ClassReader) (*ClassFile, error) {
	val := reader.ReadUInt32()
	if val != 0xCAFEBABE {
		panic("java.lang.ClassFormatCheckError: magic!")
	}
	self.magic = val
	return self, nil
}

func (self *ClassFile) ReadMinorVersion(reader *reader.ClassReader) (*ClassFile, error) {
	val := reader.ReadUInt16()
	self.majorVersion = val
	return self, nil
}

func (self *ClassFile) ReadMajorVersion(reader *reader.ClassReader) (*ClassFile, error) {
	val := reader.ReadUInt16()
	self.majorVersion = val
	return self, nil
}

func (self *ClassFile) ReadConstantPoolCount(reader *reader.ClassReader) (*ClassFile, error) {
	val := reader.ReadUInt16()
	self.constantPoolCount = val
	return self, nil
}

func (self *ClassFile) ReadAccessFlags(reader *reader.ClassReader) (*ClassFile, error) {
	val := reader.ReadUInt16()
	self.accessFlags = val
	return self, nil
}

func (self *ClassFile) ReadThisClass(reader *reader.ClassReader) (*ClassFile, error) {
	val := reader.ReadUInt16()
	self.thisClass = val
	return self, nil
}

func (self *ClassFile) ReadSuperClass(reader *reader.ClassReader) (*ClassFile, error) {
	val := reader.ReadUInt16()
	self.superClass = val
	return self, nil
}

func (self *ClassFile) ReadInterfacesCount(reader *reader.ClassReader) (*ClassFile, error) {
	val := reader.ReadUInt16()
	self.interfacesCount = val
	return self, nil
}

func (self *ClassFile) ReadFieldsCount(reader *reader.ClassReader) (*ClassFile, error) {
	val := reader.ReadUInt16()
	self.fieldsCount = val
	return self, nil
}

func (self *ClassFile) ReadAttributesCount(reader *reader.ClassReader) (*ClassFile, error) {
	val := reader.ReadUInt16()
	self.attributesCount = val
	return self, nil
}

func (self *ClassFile) Magic() uint32 {
	return self.magic
}
