package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProd        bool
	Session       *scs.SessionManager
}

func GetPath() string {
	return "/Users/neha/Go_Vacay/"
}
