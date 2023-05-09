package helpers

import (
	"Go_Vacay/internal/config"
	"fmt"
	"net/http"
	"runtime/debug"
)

var app *config.AppConfig

// set up app config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a

}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client Error with status: ", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
