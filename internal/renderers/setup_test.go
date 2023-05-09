package renderers

import (
	"Go_Vacay/internal/config"
	"Go_Vacay/internal/models"
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager
var testApp config.AppConfig

/*
*	create an interface similar to http response writer interface
*	A ResponseWriter interface is used by an HTTP handler to construct an HTTP response.
 */
type MyWriter struct{}

func (tw *MyWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *MyWriter) WriteHeader(i int) {}

func (tw *MyWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}

func TestMain(m *testing.M) {

	gob.Register(models.Reservation{})
	gob.Register(models.Registration{})
	gob.Register(models.Login{})

	//set to true if production environment
	testApp.InProd = false

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	testApp.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	testApp.ErrorLog = errorLog

	//initialise session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session
	app = &testApp

	os.Exit(m.Run())
}
