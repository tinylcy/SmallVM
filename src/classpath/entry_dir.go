package classpath

import (
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

type DirEntry struct {
	path string // directory absolute path
}

func (self *DirEntry) ReadClass(className string) ([]byte, Entry, error) {
	data, err := ioutil.ReadFile(filepath.Join(self.path, className))
	if err == nil {
		return data, self, nil
	}

	files, err := ioutil.ReadDir(self.path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileName := file.Name()
		if strings.HasSuffix(fileName, ".jar") {
			// fmt.Printf("fileName: %s\n", filepath.Join(self.path, fileName))
			zip := NewZipEntry(filepath.Join(self.path, fileName))
			data, _, err := zip.ReadClass(className)
			if err == nil {
				return data, self, nil
			}
		}
	}

	return nil, nil, errors.New("Class Not Found: " + className)
}

func NewDirEntry(path string) *DirEntry {
	return &DirEntry{path}
}
