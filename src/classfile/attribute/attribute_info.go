package attribute

import (
	"SmallVM/classfile/constantpool"
	"SmallVM/classfile/reader"
)

type AttributeInfo interface {
	readInfo(reader *reader.ClassReader)
}

func ReadAttributes(reader *reader.ClassReader, pool []constantpool.ConstantInfo, count uint16) []AttributeInfo {
	attributes := make([]AttributeInfo, 0)
	for index := 0; index < int(count); index++ {
		nameIndex := reader.ReadUInt16()
		name := pool[nameIndex-1].String()
		attribute := NewAttributeInfo(nameIndex, name, pool)
		attribute.readInfo(reader)
		attributes = append(attributes, attribute)
	}

	return attributes
}

func NewAttributeInfo(nameIndex uint16, name string, pool []constantpool.ConstantInfo) AttributeInfo {
	switch name {
	case "Code":
		return NewCode(pool, nameIndex)
	case "Exceptions":
		return NewExceptions(nameIndex)
	case "LineNumberTable":
		return NewLineNumberTable(nameIndex)
	case "LocalVariableTable":
		return NewLocalVariableTable(nameIndex)
	case "LocalVariableTypeTable":
		return NewLocalVariableTypeTable(nameIndex)
	case "SourceFile":
		return NewSourceFile(nameIndex)
	case "ConstantValue":
		return NewConstantValue(nameIndex)
	case "InnerClasses":
		return NewInnerClasses(nameIndex)
	case "Deprecated":
		return NewDeprecated(nameIndex)
	case "Synthetic":
		return NewSynthetic(nameIndex)
	case "Signature":
		return NewSignature(nameIndex)
	case "BootstrapMethods":
		return NewBootstrapMethods(nameIndex)
	default: // case "StackMapTable"
		// fmt.Printf("Get Unparsed Attribute: name = %s\n", name)
		return NewUnparsed(nameIndex)
	}

}
