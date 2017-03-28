package accessflags

import (
	"strings"
)

var CLASS_ACCESS_FLAGS = map[uint16]string{
	0x0001: "public",
	0x0010: "final",
	0x0020: "super",
	0x0200: "interface",
	0x0400: "abstract",
	0x1000: "synthetic",
	0x2000: "annotation",
	0x4000: "enum",
}

var FIELD_ACCESS_FLAGS = map[uint16]string{
	0x0001: "public",
	0x0002: "private",
	0x0004: "protected",
	0x0008: "static",
	0x0010: "final",
	0x0040: "volatile",
	0x0080: "transient",
	0x1000: "synthetic",
	0x4000: "enum",
}

var METHOD_ACCESS_FLAGS = map[uint16]string{
	0x0001: "public",
	0x0002: "private",
	0x0004: "protected",
	0x0008: "static",
	0x0010: "final",
	0x0020: "synchronized",
	0x0040: "bridge",
	0x0080: "varargs",
	0x0100: "native",
	0x0400: "abstract",
	0x0800: "strictfp",
	0x1000: "synthetic",
}

func GetClassAccessFlags(flags uint16) string {
	output := make([]string, 0)
	for k, v := range CLASS_ACCESS_FLAGS {
		if flags&k != 0 {
			output = append(output, v)
		}
	}
	return strings.Join(output, ", ")
}

func GetFieldAccessFlags(flags uint16) string {
	output := make([]string, 0)
	for k, v := range FIELD_ACCESS_FLAGS {
		if flags&k != 0 {
			output = append(output, v)
		}
	}
	return strings.Join(output, ", ")
}

func GetMethodAccessFlags(flags uint16) string {
	output := make([]string, 0)
	for k, v := range METHOD_ACCESS_FLAGS {
		if flags&k != 0 {
			output = append(output, v)
		}
	}
	return strings.Join(output, ", ")
}
