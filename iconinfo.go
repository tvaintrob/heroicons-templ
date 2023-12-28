package main

import (
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type IconInfo struct {
	Size    int
	Name    string
	Group   string
	Content string
}

func NewIconInfo(path string) IconInfo {
	re := regexp.MustCompile(`upstream/src/(.*?)/(.*?)/(.*?)\.svg`)
	match := re.FindStringSubmatch(path)

	size, group, name := match[1], match[2], match[3]
	sizenum, err := strconv.Atoi(size)
	if err != nil {
		panic(err)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return IconInfo{sizenum, buildIconName(name), group, string(content)}
}

func buildIconName(name string) string {
	words := strings.Split(name, "-")
	caser := cases.Title(language.English)
	for i, word := range words {
		words[i] = caser.String(word)
	}

	return strings.Join(words, "")
}

func (i IconInfo) GenerateTemplComponent(w io.Writer) error {
	templ := template.Must(template.New(i.Name).Parse(`
package {{.Group}}{{.Size}}

templ {{.Name}}Icon() {
  {{.Content}}
}
  `))

	return templ.Execute(w, i)
}
