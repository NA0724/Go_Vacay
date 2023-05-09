package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

var path = "/Users/neharaj/Go_Vacay/"

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProd        bool
	Session       *scs.SessionManager
}

func GetDirPath() string {
	return "/Users/neha/Go_Vacay/"
}
