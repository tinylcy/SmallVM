package classfile

import "SmallVM/classfile/reader"
import "SmallVM/classfile/constantpool"
import "SmallVM/classfile/attribute"
import "fmt"

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
	attributes        []attribute.AttributeInfo
}

type FieldInfo struct {
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributesCount uint16
	attributes      []attribute.AttributeInfo
}

type MethodInfo struct {
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributesCount uint16
	attributes      []attribute.AttributeInfo
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
	fmt.Printf("constantpool length = %d\n", self.constantPoolCount)

	self.ReadAccessFlags(reader)
	self.ReadThisClass(reader)
	self.ReadSuperClass(reader)

	self.ReadInterfacesCount(reader)
	fmt.Printf("interfacesCount = %d\n", self.interfacesCount)
	self.ReadInterfaces(reader, self.interfacesCount)

	self.ReadFieldsCount(reader)
	fmt.Printf("fieldsCount = %d\n", self.fieldsCount)
	self.ParseFields(reader, self.fieldsCount)

	self.ReadMethodsCount(reader)
	fmt.Printf("methodsCount = %d\n", self.methodsCount)
	self.ParseMethods(reader, self.methodsCount)

	self.ReadAttributesCount(reader)
	fmt.Printf("class attributes count = %d\n", self.attributesCount)
	self.ParseClassAttributes(reader, self.attributesCount)

	return self, nil
}

func (self *ClassFile) ParseFields(reader *reader.ClassReader, count uint16) (*ClassFile, error) {
	self.fields = make([]FieldInfo, count)
	for index := 0; index < int(count); index++ {
		field := FieldInfo{}
		field.accessFlags = reader.ReadUInt16()
		field.nameIndex = reader.ReadUInt16()
		fmt.Printf("field nameIndex = %d\n", field.nameIndex)
		field.descriptorIndex = reader.ReadUInt16()
		field.attributesCount = reader.ReadUInt16()
		fmt.Printf("field attributes count = %d\n", field.attributesCount)
		field.attributes = attribute.ReadAttributes(reader, self.constantPool, field.attributesCount)
		self.fields = append(self.fields, field)
	}
	return self, nil
}

func (self *ClassFile) ParseMethods(reader *reader.ClassReader, count uint16) (*ClassFile, error) {
	self.methods = make([]MethodInfo, 0)
	for index := 0; index < int(count); index++ {
		method := MethodInfo{}
		method.accessFlags = reader.ReadUInt16()
		method.nameIndex = reader.ReadUInt16()
		fmt.Printf("method nameIndex = %d\n", method.nameIndex)
		method.descriptorIndex = reader.ReadUInt16()
		method.attributesCount = reader.ReadUInt16()
		fmt.Printf("method attributes count = %d\n", method.attributesCount)
		method.attributes = attribute.ReadAttributes(reader, self.constantPool, method.attributesCount)
		self.methods = append(self.methods, method)
	}
	return self, nil
}

// Parse class level attributes.
func (self *ClassFile) ParseClassAttributes(reader *reader.ClassReader, count uint16) (*ClassFile, error) {
	self.attributes = attribute.ReadAttributes(reader, self.constantPool, count)
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

func (self *ClassFile) ReadInterfaces(reader *reader.ClassReader, count uint16) (*ClassFile, error) {
	self.interfaces = make([]uint16, count)
	for index := 0; index < int(count); index++ {
		self.interfaces = append(self.interfaces, reader.ReadUInt16())
	}
	return self, nil
}

func (self *ClassFile) ReadFieldsCount(reader *reader.ClassReader) (*ClassFile, error) {
	val := reader.ReadUInt16()
	self.fieldsCount = val
	return self, nil
}

func (self *ClassFile) ReadMethodsCount(reader *reader.ClassReader) (*ClassFile, error) {
	val := reader.ReadUInt16()
	self.methodsCount = val
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

func (self *ClassFile) InterfacesCount() uint16 {
	return self.interfacesCount
}

func (self *ClassFile) Interfaces() []uint16 {
	return self.interfaces
}
