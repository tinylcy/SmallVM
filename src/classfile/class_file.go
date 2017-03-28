package classfile

import (
	"SmallVM/classfile/accessflags"
	"SmallVM/classfile/attribute"
	"SmallVM/classfile/constantpool"
	"SmallVM/classfile/reader"
	"fmt"
	"strconv"
	"strings"
)

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

	// parse constant pool info
	self.ReadConstantPoolCount(reader)
	self.constantPool = constantpool.Parse(reader, self.constantPoolCount)

	self.ReadAccessFlags(reader)
	self.ReadThisClass(reader)
	self.ReadSuperClass(reader)

	self.ReadInterfacesCount(reader)
	self.ReadInterfaces(reader, self.interfacesCount)

	// parse fields info
	self.ReadFieldsCount(reader)
	self.ParseFields(reader, self.fieldsCount)

	// parse methods info
	self.ReadMethodsCount(reader)
	self.ParseMethods(reader, self.methodsCount)

	// parse attributes info
	self.ReadAttributesCount(reader)
	self.ParseClassAttributes(reader, self.attributesCount)

	return self, nil
}

func (self *ClassFile) ParseFields(reader *reader.ClassReader, count uint16) (*ClassFile, error) {
	self.fields = make([]FieldInfo, 0)
	for index := 0; index < int(count); index++ {
		field := FieldInfo{}
		field.accessFlags = reader.ReadUInt16()
		field.nameIndex = reader.ReadUInt16()
		field.descriptorIndex = reader.ReadUInt16()
		field.attributesCount = reader.ReadUInt16()
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
		method.descriptorIndex = reader.ReadUInt16()
		method.attributesCount = reader.ReadUInt16()
		method.attributes = attribute.ReadAttributes(reader, self.constantPool, method.attributesCount)
		self.methods = append(self.methods, method)
	}
	return self, nil
}

// Print class structure
func (self *ClassFile) ViewClass() {
	fmt.Printf("magic: %x\n", self.magic)
	fmt.Printf("minor version: %d\n", self.minorVersion)
	fmt.Printf("major version: %d\n", self.majorVersion)
	fmt.Printf("constant pool count: %d\n", self.constantPoolCount)
	fmt.Printf("access flags: %s\n", self.AccessFlags(self.accessFlags))
	fmt.Printf("this class: %s\n", self.ThisClass())
	fmt.Printf("super class: %s\n", self.SuperClass())
	fmt.Printf("interfaces: %s\n", self.Interfaces())
	fmt.Printf("fields count: %d\n", self.fieldsCount)
	self.ViewFields()
	fmt.Printf("method count: %d\n", self.methodsCount)
	self.ViewMethods()
}

// Print class fields info
func (self *ClassFile) ViewFields() {
	fields := self.fields
	for _, field := range fields {
		fmt.Println()
		fmt.Printf("field name: %s\n", self.getUtf8Value(field.nameIndex))
		fmt.Printf("access flags: %s\n", accessflags.GetFieldAccessFlags(field.accessFlags))
		fmt.Printf("descriptor index: %s\n", self.getUtf8Value(field.descriptorIndex))
	}
	fmt.Println()
}

// Print class methods info
func (self *ClassFile) ViewMethods() {
	methods := self.methods
	for _, method := range methods {
		fmt.Println()
		fmt.Printf("method name: %s\n", self.getUtf8Value(method.nameIndex))
		fmt.Printf("access flags: %s\n", accessflags.GetMethodAccessFlags(method.accessFlags))
		fmt.Printf("descriptor index: %s\n", self.getUtf8Value(method.descriptorIndex))
	}
}

// Get constantpool's CONSTANT_Utf8_info value.
func (self *ClassFile) getUtf8Value(index uint16) string {
	if index <= 0 || index >= self.constantPoolCount {
		panic("GetUtf8Value: Index Out of Bounds")
	}
	val := self.constantPool[index-1].String()
	return val
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
	var index uint16
	self.interfaces = make([]uint16, 0)
	for index = 0; index < count; index++ {
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

func (self *ClassFile) AccessFlags(flags uint16) string {
	return accessflags.GetClassAccessFlags(flags)
}

func (self *ClassFile) ThisClass() string {
	index, _ := strconv.Atoi(self.getUtf8Value(self.thisClass))
	return self.getUtf8Value(uint16(index))
}

func (self *ClassFile) SuperClass() string {
	index, _ := strconv.Atoi(self.getUtf8Value(self.superClass))
	return self.getUtf8Value(uint16(index))
}

func (self *ClassFile) Interfaces() string {
	var index int
	interfaces := make([]string, 0)
	for index = 0; index < len(self.interfaces); index++ {
		interIndex, _ := strconv.Atoi(self.getUtf8Value(self.interfaces[index]))
		interfaces = append(interfaces, self.getUtf8Value(uint16(interIndex)))
	}
	return strings.Join(interfaces, ", ")
}
