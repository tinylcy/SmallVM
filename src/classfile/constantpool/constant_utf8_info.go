package constantpool

import "SmallVM/classfile/reader"
import "fmt"

type ConstantUtf8Info struct {
	val string
}

func NewConstantUtf8Info() *ConstantUtf8Info {
	return &ConstantUtf8Info{}
}

func (self *ConstantUtf8Info) readInfo(reader *reader.ClassReader) {
	length := reader.ReadUInt16()
	bytes := reader.ReadBytes(uint32(length))
	self.val = string(bytes)
	fmt.Printf("ConstantUtf8Info: %s\n", self.val)
}

// Convert MUTF-8 byte slice to UTF-8 byte array
func (self *ConstantUtf8Info) mutf8ToString(bytes []byte) string {
	var nilCount int = 0
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == 0 {
			nilCount++
		}
	}

	convertedBytes := make([]byte, len(bytes)+nilCount)

	for i, j := 0, 0; i < len(bytes); {
		convertedBytes[j] = bytes[i]
		if bytes[i] == 0 {
			convertedBytes[j] = byte(0xC0)
			j++
			convertedBytes[j] = byte(0x80)
		}
		i++
		j++
	}

	return string(convertedBytes)
}
