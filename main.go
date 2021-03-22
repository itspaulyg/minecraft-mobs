package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

// Generic to each html page
type Page struct {
	Title string
	Mob   string
}

type MobType int

const (
	Unknown = iota
	Passive
	Hostile
)

// Keep track of html files
var templates = template.Must(template.ParseGlob("templates/*.html"))
var passive_templates = template.Must(template.ParseGlob("templates/passive/*.html"))
var hostile_templates = template.Must(template.ParseGlob("templates/hostile/*.html"))

// Get port either local or heroku hosted
func getPort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return ":" + port
	}
	return ":8080"
}

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	// Home
	http.HandleFunc("/", Home)

	// Passive Mobs
	passiveMobs := []string{
		"chicken",
		"cow",
		"pig",
		"sheep",
		"villager",
	}
	for _, mob := range passiveMobs {
		http.HandleFunc("/"+strings.ReplaceAll(mob, " ", "-")+"/", mobHandler(mob, Passive))
	}

	// Hostile Mobs
	hostileMobs := []string{
		"blaze",
		"creeper",
		"ghast",
		"magma cube",
		"skeleton",
		"slime",
		"zombie",
	}
	for _, mob := range hostileMobs {
		http.HandleFunc("/"+strings.ReplaceAll(mob, " ", "-")+"/", mobHandler(mob, Hostile))
	}

	log.Fatal(http.ListenAndServe(getPort(), nil))
}

// Render the template to the view
func renderTemplate(w http.ResponseWriter, tmpl string, p Page, check int) {
	if check == 1 {
		err := passive_templates.ExecuteTemplate(w, tmpl+".html", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if check == 2 {
		err := hostile_templates.ExecuteTemplate(w, tmpl+".html", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		err := templates.ExecuteTemplate(w, tmpl+".html", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func mobHandler(mob string, mobType MobType) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		p := Page{
			Title: "Mobs: " + strings.Title(mob),
			Mob:   strings.Title(mob),
		}
		renderTemplate(w, strings.Join(strings.Fields(strings.ToLower(mob)), ""), p, int(mobType))
	}
}

// Render Home Template
func Home(w http.ResponseWriter, r *http.Request) {
	p := Page{
		Title: "Minecraft Mobs",
	}
	renderTemplate(w, "index", p, 0)
}
