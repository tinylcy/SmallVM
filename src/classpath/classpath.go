package classpath

import (
	"SmallVM/cmd"
	"errors"
	"path/filepath"
	"strings"
)

type Classpath struct {
	bootstrap   Entry // bootstrap classpath
	extension   Entry // extension classpath
	application Entry // application classpath
}

func Parse(cmd *cmd.CmdLine) *Classpath {
	cp := &Classpath{}

	// find JRE path
	jrePath, err := cmd.JREPath()
	if err != nil {
		panic("JRE path does not exist!")
	}
	// fmt.Printf("JRE path: %s\n", jrePath)

	jreLibPath := filepath.Join(jrePath, "lib")
	cp.bootstrap = NewDirEntry(jreLibPath)
	// fmt.Printf("bootstrap classpath: %s\n", jreLibPath)

	jreExtPath := filepath.Join(jrePath, "lib", "ext")
	cp.extension = NewDirEntry(jreExtPath)
	// fmt.Printf("extension classpath: %s\n", jreExtPath)

	appClassPath, _ := cmd.AppClassPath()
	cp.application = NewDirEntry(appClassPath)
	// fmt.Printf("application classpath: %s\n", appClassPath)

	return cp
}

func (self *Classpath) ReadClass(className string) ([]byte, error) {
	className = strings.Replace(className, ".", "/", -1)
	className = className + ".class"

	if data, _, err := self.bootstrap.ReadClass(className); err == nil {
		return data, nil
	}
	if data, _, err := self.extension.ReadClass(className); err == nil {
		return data, nil
	}

	if data, _, err := self.application.ReadClass(className); err == nil {
		return data, nil
	}

	return nil, errors.New("Class Not Found: " + className)
}
