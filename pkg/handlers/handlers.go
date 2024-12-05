package handlers

import (
	"net"
	"net/http"

	"github.com/abhinsyam/bookings/pkg/config"
	"github.com/abhinsyam/bookings/pkg/models"
	"github.com/abhinsyam/bookings/pkg/render"
)

// Repo is the repository used by handlers
var Repo *Repository

// Repository type
type Repository struct {
	App *config.AppConfig
}

// NEW REPO sets the repository for the handlers
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// New handlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		http.Error(w, "Unable to parse remote address", http.StatusInternalServerError)
		return
	}

	// Handle IPv6 loopback (::1) and convert it to IPv4 loopback (127.0.0.1)
	if remoteIP == "::1" {
		remoteIP = "127.0.0.1"
	}

	//remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
