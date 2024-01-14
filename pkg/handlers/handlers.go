package handlers

import (
	"net/http"

	"github.com/thiruthanikaiarasu/web/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderedTemplate(w, "home.page.html")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderedTemplate(w, "about.page.html")
}

