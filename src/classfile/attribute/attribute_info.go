package attribute

import "SmallVM/classfile/reader"
import "SmallVM/classfile/constantpool"
import "fmt"

type AttributeInfo interface {
	readInfo(reader *reader.ClassReader)
}

func ReadAttributes(reader *reader.ClassReader, pool []constantpool.ConstantInfo, count uint16) []AttributeInfo {
	attributes := make([]AttributeInfo, 0)
	for index := 0; index < int(count); index++ {
		nameIndex := reader.ReadUInt16()
		name := pool[nameIndex-1].String()
		fmt.Printf("attribute name = %s\n", name)
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
		return NewExceptions()
	case "LineNumberTable":
		return NewLineNumberTable(nameIndex)
	case "LocalVariableTable":
		return NewLocalVariableTable(nameIndex)
	case "SourceFile":
		return NewSourceFile(nameIndex)
	case "ConstantValue":
		return NewConstantValue(nameIndex)
	case "InnerClasses":
		return NewInnerClasses()
	case "Deprecated":
		return NewDeprecated()
	case "Synthetic":
		return NewSynthetic()
	case "StackMapTable":
		return NewStackMapTable()
	case "Signature":
		return NewSignature()
	case "BootstrapMethods":
		return NewBootstrapMethods()
	default:
		panic("No Such Attribute: " + name)
	}

}
