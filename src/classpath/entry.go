package classpath

import (
	"os"
	"path/filepath"
)

const PathSeperator = filepath.Separator
const PathListSeparator = string(os.PathListSeparator)

type Entry interface {
	ReadClass(className string) ([]byte, Entry, error)
}
