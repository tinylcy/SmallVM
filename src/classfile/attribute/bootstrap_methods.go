package attribute

import "SmallVM/classfile/reader"

type BootstrapMethods struct {
	attributeNameIndex  uint16
	attributeLength     uint32
	numBootstrapMethods uint16
	bootstrapMethods    []BootstrapMethod
}

type BootstrapMethod struct {
	bootstrapMethodRef    uint16
	numBootstrapArguments uint16
	bootstrapArguments    []uint16
}

func NewBootstrapMethods(nameIndex uint16) *BootstrapMethods {
	return &BootstrapMethods{attributeNameIndex: nameIndex}
}

func (self *BootstrapMethods) readInfo(reader *reader.ClassReader) {
	self.attributeLength = reader.ReadUInt32()
	self.numBootstrapMethods = reader.ReadUInt16()
	self.bootstrapMethods = self.readBootstrapMethods(reader, self.numBootstrapMethods)
}

func (self *BootstrapMethods) readBootstrapMethods(reader *reader.ClassReader, num uint16) []BootstrapMethod {
	var methodCnt, augCnt uint16 = 0, 0
	methods := make([]BootstrapMethod, 0)
	for methodCnt = 0; methodCnt < num; methodCnt++ {
		method := BootstrapMethod{}
		method.bootstrapMethodRef = reader.ReadUInt16()
		method.numBootstrapArguments = reader.ReadUInt16()
		method.bootstrapArguments = make([]uint16, 0)
		for augCnt = 0; augCnt < method.numBootstrapArguments; augCnt++ {
			method.bootstrapArguments = append(method.bootstrapArguments, reader.ReadUInt16())
		}
		methods = append(methods, method)
	}
	return methods
}
