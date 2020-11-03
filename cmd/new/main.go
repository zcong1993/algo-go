package new

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/zcong1993/algo-go/pkg/leetcode"
)

const (
	folder = "solve"
	prefix = "solve"
)

var (
	codeTpl    = template.Must(template.New("code").Parse(codeStr))
	problemTpl = template.Must(template.New("problem").Parse(problemStr))
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func normalizeNumber(number string) string {
	if len(number) >= 4 {
		return number
	}
	return strings.Repeat("0", 4-len(number)) + number
}

type MetaWithFolder struct {
	leetcode.Meta
	Folder string
	TagStr string
}

func Run(n string) {
	meta, err := leetcode.GetMetaByNumber(n)
	if err != nil || meta == nil {
		log.Fatal(err, meta)
	}
	number := meta.Index
	folderName := prefix + normalizeNumber(number)
	fp := filepath.Join(folder, folderName)
	if fileExists(fp) {
		log.Fatalf("path %s exists", fp)
	}
	os.MkdirAll(fp, 0755)
	codeFp := filepath.Join(fp, fmt.Sprintf("solve_%s.go", number))
	codeTestFp := filepath.Join(fp, fmt.Sprintf("solve_%s_test.go", number))
	problemFp := filepath.Join(fp, "problem.md")
	metaf := &MetaWithFolder{
		*meta,
		folderName,
		strings.Join(meta.Tags, ","),
	}
	metaf.Meta.Content = strings.ReplaceAll(metaf.Meta.Content, "↵", "")
	var codeContent bytes.Buffer
	err = codeTpl.Execute(&codeContent, metaf)
	if err != nil {
		log.Fatal(err)
	}
	if !fileExists(codeFp) {
		ioutil.WriteFile(codeFp, codeContent.Bytes(), 0644)
	}
	if !fileExists(codeTestFp) {
		ioutil.WriteFile(codeTestFp, []byte(fmt.Sprintf(testCodeStr, folderName)), 0644)
	}
	var problemContent bytes.Buffer
	err = problemTpl.Execute(&problemContent, metaf)
	if err != nil {
		log.Fatal(err)
	}
	if !fileExists(problemFp) {
		ioutil.WriteFile(problemFp, problemContent.Bytes(), 0644)
	}
}

var codeStr = `package {{ .Folder }}

/**
@index {{ .Index }}
@title {{ .Title }}
@difficulty {{ .Difficulty }}
@tags {{ .TagStr }}
@draft false
@link {{ .Link }}
*/
func solve() {

}
`

var testCodeStr = `package %s

`

var problemStr = `# [{{ .Index }}. {{ .Title }}]({{ .Link }})

{{ .Content }}
`
