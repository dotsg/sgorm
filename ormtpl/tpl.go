package ormtpl

import (
	"embed"
	"text/template"
)

//go:embed */*.gogo
var templates embed.FS
var tpl = template.New("tpl")

// GetTpl return Template of giving path
func GetTpl(driver, path string) *template.Template {
	var err error

	tplStr, err := templates.ReadFile(driver + "/" + path)
	if err != nil {
		tplStr, _ = templates.ReadFile("default/" + path)
	}

	result, err := tpl.Parse(string(tplStr))
	if err != nil {
		panic(err)
	}
	return result
}

// TplStruct define object used for tpl code generation
type TplStruct interface {
	SetVersion(version string)
	GetVersion() string
	GetName() string
}
