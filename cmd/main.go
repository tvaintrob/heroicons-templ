package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

func main() {
	icons := findIcons()
	log.Printf("generating %d icons...", len(icons))

	for _, icon := range icons {
		pkgpath := path.Join(icon.Group, strconv.FormatInt(int64(icon.Size), 10))
		filepath := path.Join(pkgpath, fmt.Sprintf("%s.templ", icon.Name))

		if err := os.MkdirAll(pkgpath, 0750); err != nil {
			log.Fatal(err)
		}

		w, err := os.Create(filepath)
		if err != nil {
			log.Fatal(err)
		}

		if err := icon.GenerateTemplComponent(w); err != nil {
			log.Fatal(err)
		}
	}
}

func findIcons() []IconInfo {
	var icons []IconInfo
	_ = filepath.WalkDir("upstream/src", func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".svg" {
			icons = append(icons, NewIconInfo(path))
		}
		return nil
	})

	return icons
}
