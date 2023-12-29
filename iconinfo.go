package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/beevik/etree"
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

	info := IconInfo{sizenum, buildIconName(name), group, string(content)}
	switch sizenum {
	case 24:
		info.fixSvgContent("w-6 h-6")
	case 20:
		info.fixSvgContent("w-5 h-5")
	case 16:
		info.fixSvgContent("w-4 h-4")
	}

	return info
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

func (i *IconInfo) fixSvgContent(classes string) error {
	doc := etree.NewDocument()
	if err := doc.ReadFromString(i.Content); err != nil {
		return err
	}

	element := doc.FindElement("//svg")
	if element == nil {
		return fmt.Errorf("no svg element found")
	}

	element.CreateAttr("class", classes)
	element.CreateAttr("fill", "currentColor")

	for _, e := range element.ChildElements() {
		e.RemoveAttr("fill")
	}

	updatedContent, err := doc.WriteToString()
	if err != nil {
		return err
	}

	i.Content = updatedContent
	return nil
}
