package handlers

import (
	"code-database/config"
	"code-database/development"
	"html/template"
)

func getTemplate(filePaths ...string) (*template.Template, error) {
	if config.BuildMode == "dev" {
		return development.GetTemplateFromDevServer(filePaths...)
	} else {
		return template.ParseFiles(filePaths...)
	}
}

func getTemplateWithFuncs(funcMap template.FuncMap, filePaths ...string) (*template.Template, error) {
	if config.BuildMode == "dev" {
		return development.GetTemplateFromDevServerWithFuncs(funcMap, filePaths...)
	} else {
		return template.ParseFiles(filePaths...)
	}
}
