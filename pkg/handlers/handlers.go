package handlers

import (
	//"errors"
	// "fmt"
	// "html/template"

	"github.com/scottywm/bookings/pkg/config"
	"github.com/scottywm/bookings/pkg/models"
	"github.com/scottywm/bookings/pkg/render"
	"net/http"
)

// Repo .... The repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr                              // we get the ip address for each user
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP) // we store the ip address inside the Session property in the Repo variable for each request. The key is titled "remote_ip" and its valuse is the value of the remoteIP variable.
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	// here we are getting the ip address that was stored inside the Session property.
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
