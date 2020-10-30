package new

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	folder = "solve"
	prefix = "solve"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func Run(number string) {
	folderName := prefix + number
	fp := filepath.Join(folder, folderName)
	if fileExists(fp) {
		log.Fatalf("path %s exists", fp)
	}
	os.MkdirAll(fp, 0755)
	codeFp := filepath.Join(fp, fmt.Sprintf("solve_%s.go", number))
	codeTestFp := filepath.Join(fp, fmt.Sprintf("solve_%s_test.go", number))
	if !fileExists(codeFp) {
		ioutil.WriteFile(codeFp, []byte(fmt.Sprintf(codeStr, folderName)), 0644)
	}

	if !fileExists(codeTestFp) {
		ioutil.WriteFile(codeTestFp, []byte(fmt.Sprintf(testCodeStr, folderName)), 0644)
	}
}

var codeStr = `package %s

/**
@index
@title
@difficulty 简单
@tags normal
@draft false
@link
*/
func solve() {

}
`

var testCodeStr = `package %s

`
