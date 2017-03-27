package constantpool

import "SmallVM/classfile/reader"

const (
	CONSTANT_UTF8_INFO            = 1
	CONSTANT_INTEGER_INFO         = 3
	CONSTANT_FLOAT_INFO           = 4
	CONSTANT_LONG_INFO            = 5
	CONSTANT_DOUBLE_INFO          = 6
	CONSTANT_CLASS_INFO           = 7
	CONSTANT_STRING_INGFO         = 8
	CONSTANT_FIELD_REF_INFO       = 9
	CONSTANT_METHOD_REF_INFO      = 10
	CONSTANT_INTERFACE_METHOD_REF = 11
	CONSTANT_NAME_AND_TYPE_INFO   = 12
	CONSTANT_METHOD_HANDLE_INFO   = 15
	CONSTANT_METHOD_TYPE_TYPE     = 16
	CONSTANT_INVOKE_DYNAMIC_INFO  = 18
)

type ConstantInfo interface {
	readInfo(reader *reader.ClassReader)
	String() string
}

// Parse constant pool, return a slice of contant info.
func Parse(reader *reader.ClassReader, constantPoolCount uint16) []ConstantInfo {
	infoList := make([]ConstantInfo, 0)
	for index := 1; index < int(constantPoolCount); index++ {
		tag := reader.ReadUInt8()
		if tag == 0 {
			reader.ReadUInt8() // read one byte when tag == 0
			continue           // tag == 0 means don't refer to any constant
		}
		info := NewConstantInfo(tag)
		info.readInfo(reader)
		infoList = append(infoList, info)
	}
	// fmt.Printf("infoList length = %d\n", len(infoList))
	return infoList
}

func NewConstantInfo(tag uint8) ConstantInfo {
	switch tag {
	case CONSTANT_UTF8_INFO:
		return NewConstantUtf8Info()
	case CONSTANT_INTEGER_INFO:
		return NewConstantIntegetInfo()
	case CONSTANT_FLOAT_INFO:
		return NewConstantFloatInfo()
	case CONSTANT_LONG_INFO:
		return NewConstantLongInfo()
	case CONSTANT_DOUBLE_INFO:
		return NewConstantDoubleInfo()
	case CONSTANT_CLASS_INFO:
		return NewConstantClassInfo()
	case CONSTANT_STRING_INGFO:
		return NewConstantStringInfo()
	case CONSTANT_FIELD_REF_INFO:
		return NewConstantFieldRefInfo()
	case CONSTANT_METHOD_REF_INFO:
		return NewConstantMethodRefInfo()
	case CONSTANT_INTERFACE_METHOD_REF:
		return NewConstantInterfaceMethodRefInfo()
	case CONSTANT_NAME_AND_TYPE_INFO:
		return NewConstantNameAndTypeInfo()
	case CONSTANT_METHOD_HANDLE_INFO:
		return NewConstantMethodHandleInfo()
	case CONSTANT_METHOD_TYPE_TYPE:
		return NewConstantMethodTypeInfo()
	case CONSTANT_INVOKE_DYNAMIC_INFO:
		return NewConstantInvokeDynamicInfo()
	default:
		panic("Cannot Find ConstantInfo, tag: " + string(tag))
	}
}
