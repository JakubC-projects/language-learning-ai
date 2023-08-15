package router

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func loadTemplates(r *gin.Engine) {
	r.SetHTMLTemplate(htmlTemplate)
}

var htmlTemplate *template.Template

func init() {
	tmpl, err := findAndParseTemplates("templates", nil)
	if err != nil {
		panic(fmt.Errorf("cannot parse templates: %w", err))
	}
	htmlTemplate = tmpl
}

func findAndParseTemplates(rootDir string, funcMap template.FuncMap) (*template.Template, error) {
	cleanRoot := filepath.Clean(rootDir)
	pfx := len(cleanRoot) + 1
	root := template.New("")

	err := filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
		if !info.IsDir() {
			if e1 != nil {
				return e1
			}

			b, e2 := os.ReadFile(path)
			if e2 != nil {
				return e2
			}

			name := path[pfx:]
			t := root.New(name).Funcs(funcMap)
			_, e2 = t.Parse(string(b))
			if e2 != nil {
				return e2
			}
		}

		return nil
	})

	return root, err
}
