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

func NewBootstrapMethods() *BootstrapMethod {
	return &BootstrapMethod{}
}

func (self *BootstrapMethod) readInfo(reader *reader.ClassReader) {

}
