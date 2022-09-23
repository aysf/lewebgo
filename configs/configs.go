package configs

import "html/template"

type AppConfig struct {
	TemplateCache map[string]*template.Template
}
