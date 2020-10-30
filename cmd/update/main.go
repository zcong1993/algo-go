package update

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/bmatcuk/doublestar/v2"
)

var (
	indexRegex      = regexp.MustCompile("@index (.+)")
	titleRegex      = regexp.MustCompile("@title (.+)")
	difficultyRegex = regexp.MustCompile("@difficulty (.+)")
	tagsRegex       = regexp.MustCompile("@tags (.+)")
	draftRegex      = regexp.MustCompile("@draft (.+)")
	linkRegex       = regexp.MustCompile("@link (.+)")
)

var (
	tableTpl = template.Must(template.New("table").Parse(tableStr))
)

type Meta struct {
	Index      string
	Title      string
	Difficulty string
	Tags       []string
	Draft      bool
	Fp         string
	Link       string
}

type TagMetas map[string]([]*Meta)

func addMeta(tagMetas TagMetas, meta *Meta) {
	if meta == nil {
		return
	}
	for _, tag := range meta.Tags {
		if _, ok := tagMetas[tag]; !ok {
			tagMetas[tag] = make([]*Meta, 0)
		}
		tagMetas[tag] = append(tagMetas[tag], meta)
	}
}

func findTag(content []byte, r *regexp.Regexp) string {
	res := r.FindSubmatch(content)
	return string(res[1])
}

func findMeta(content []byte, fp string) *Meta {
	draft := findTag(content, draftRegex) == "true"
	if draft {
		return nil
	}
	tags := strings.Split(findTag(content, tagsRegex), ",")
	return &Meta{
		Index:      findTag(content, indexRegex),
		Title:      findTag(content, titleRegex),
		Difficulty: findTag(content, difficultyRegex),
		Tags:       tags,
		Draft:      draft,
		Fp:         fp,
		Link:       findTag(content, linkRegex),
	}
}

func genTable(metas []*Meta) string {
	var bf bytes.Buffer
	tableTpl.Execute(&bf, metas)
	return bf.String()
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func Run() {
	files, err := doublestar.Glob("./solve/**/**.go")
	if err != nil {
		log.Fatal(err)
	}
	tagMetas := make(TagMetas, 0)
	for _, fp := range files {
		if strings.HasSuffix(fp, "test.go") {
			continue
		}
		content, err := ioutil.ReadFile(fp)
		if err != nil {
			log.Fatal(err)
		}
		meta := findMeta(content, fp)
		addMeta(tagMetas, meta)
	}
	for tag, metas := range tagMetas {
		fp := fmt.Sprintf("%s.md", tag)
		if !fileExists(fp) {
			continue
		}
		content, err := ioutil.ReadFile(fp)
		if err != nil {
			log.Fatal(err)
		}
		table := genTable(metas)
		contents := strings.Split(string(content), "<!--- table -->")
		contents[1] = "<!--- table -->\n" + table
		newContent := strings.Join(contents, "")
		ioutil.WriteFile(fp, []byte(newContent), 0644)
	}
}

var tableStr = `
| 序号 | 难度 | 题目                    | 解答                      |
| ---- | ---- | ------------------ | ---------------- |{{ range . }}
| {{ .Index }} | {{ .Difficulty }} | [{{ .Title }}]({{ .Link }}) | [{{ .Fp }}]({{ .Fp }})|{{ end }}
`
