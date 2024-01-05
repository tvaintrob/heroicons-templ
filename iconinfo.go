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

// NewIconInfo parses an icon file to return a new instance of an IconInfo struct
// this struct is the basis of the file generation,
// in order to generate a templ component from an IconInfo use `IconInfo.generateTemplComponent`
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

// buildIconName will build the corrent icon name from the parsed file,
// the name should be a title case string of the filename
//
// Example:
//
//	name := buildIconName("arrow-up")
//	fmt.Println(name) // outputs: "ArrowUp"
func buildIconName(name string) string {
	words := strings.Split(name, "-")
	caser := cases.Title(language.English)
	for i, word := range words {
		words[i] = caser.String(word)
	}

	return strings.Join(words, "")
}

// generateTemplComponent will go over the content of the icon
// and will generate the appropriate templ component
func (i IconInfo) generateTemplComponent(w io.Writer) error {
	var class string
	switch i.Size {
	case 24:
		class = "w-6 h-6"
	case 20:
		class = "w-5 h-5"
	case 16:
		class = "w-4 h-4"
	}

	doc := etree.NewDocument()
	if err := doc.ReadFromString(i.Content); err != nil {
		return err
	}

	element := doc.FindElement("//svg")
	if element == nil {
		return fmt.Errorf("no svg element found")
	}

	for _, e := range element.ChildElements() {
		e.RemoveAttr("fill")
	}

	tdata := map[string]any{
		"Name":    i.Name,
		"Size":    i.Size,
		"Group":   i.Group,
		"Class":   class,
		"Content": getInnerHTML(element),
	}

	templ := template.Must(template.New(i.Name).Parse(`package {{.Group}}{{.Size}}

templ {{.Name}}Icon(attrs templ.Attributes) {
  <svg 
    xmlns="http://www.w3.org/2000/svg"
    fill="currentColor"
    viewBox="0 0 {{.Size}} {{.Size}}"
    class="{{.Class}}"
    { attrs... }
  >
    {{.Content}}
  </svg>
}
	`))

	return templ.Execute(w, tdata)
}

func getInnerHTML(svg *etree.Element) string {
	var innerHTML string
	children := svg.ChildElements()

	for _, child := range children {
		var builder strings.Builder
		child.WriteTo(&builder, &etree.WriteSettings{CanonicalEndTags: true})
		innerHTML += builder.String()
	}

	return innerHTML
}
