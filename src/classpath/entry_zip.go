package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"log"
)

type ZipEntry struct {
	path string
}

func (self *ZipEntry) ReadClass(className string) ([]byte, Entry, error) {
	// Open a zip archive for reading
	r, err := zip.OpenReader(self.path)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	defer r.Close()

	// Iterate through the files in the archive
	for _, f := range r.File {
		// fmt.Printf("Contents of %s:\n", f.Name)
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				log.Fatal(err)
			}

			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				log.Fatal(err)
				return nil, nil, err
			}
			return data, self, nil
		}

	}

	return nil, nil, errors.New("Class Not Found: " + className)
}

func NewZipEntry(path string) *ZipEntry {
	return &ZipEntry{path}
}
