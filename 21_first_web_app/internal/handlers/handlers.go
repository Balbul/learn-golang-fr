package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/Balbul/firstwebapp/config"
	"github.com/Balbul/firstwebapp/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	names := make(map[string]string)
	names["admin"] = "B4LBU"

	renderTemplate(w, "home", &models.TemplateData{
		StringData: names,
	})
}

func Contact(w http.ResponseWriter, r *http.Request) {

	ages := make(map[string]int)
	ages["age"] = 28
	renderTemplate(w, "contact", &models.TemplateData{
		IntData: ages,
	})
}

var appConfig *config.Config

func CreateTemplate(app *config.Config) {
	appConfig = app
}

func renderTemplate(w http.ResponseWriter, tmplName string, td *models.TemplateData) {
	templateCache := appConfig.TemplateCache

	tmpl, ok := templateCache[tmplName+".page.tmpl"]
	if !ok {
		http.Error(w, fmt.Sprintf("template %s not found", tmplName), http.StatusNotFound)
		return
	}

	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, td)
	buffer.WriteTo(w)
}

// createTemplateCache is a function that creates a cache of HTML templates.
// It scans the templates directory for .page.tmpl files and their corresponding layouts.
// The function returns a map of template names to their parsed templates and an error if any.
func CreateTemplateCache() (map[string]*template.Template, error) {
	// Initialize an empty map to store the parsed templates.
	cache := map[string]*template.Template{}

	// Glob for all .page.tmpl files in the templates directory.
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	// If there's an error, return an empty cache and the error.
	if err != nil {
		return cache, err
	}

	// Loop through each .page.tmpl file.
	for _, page := range pages {
		// Get the base name of the file (e.g., "home.page.tmpl" becomes "home.page.tmpl").
		name := filepath.Base(page)

		// Parse the .page.tmpl file into a template.
		tmpl := template.Must(template.ParseFiles(page))

		// Glob for all .layout.tmpl files in the layouts directory.
		layouts, err := filepath.Glob("./templates/layouts/*.layout.tmpl")

		// If there's an error, return the current cache and the error.
		if err != nil {
			return cache, err
		}

		// If there are any layout files, parse them into the template.
		if len(layouts) > 0 {
			tmpl.ParseGlob("./templates/layouts/*.layout.tmpl")
		}

		// Add the parsed template to the cache.
		cache[name] = tmpl
	}

	// Return the cache and nil error.
	return cache, nil
}
