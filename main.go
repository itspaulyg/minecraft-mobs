package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	
	"github.com/itspaulyg/minecraft-mobs/content"
	"github.com/itspaulyg/minecraft-mobs/model"
)

// Generic to each html page
type PageProps struct {
	Title   string
	Mob     string
	Content model.Content
}

// Keep track of html files
var templates = template.Must(template.New("main").Funcs(template.FuncMap{
	"makeFilename": func(mob string) string {
		return strings.ReplaceAll(strings.ToLower(mob), " ", "-")
	},
	"half": func(hp int) int {
		return hp/2
	},
	"hearts": func(hp int) []int {
		hearts := make([]int, hp/2)
		return hearts
	},
}).ParseGlob("templates/*.html"))

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
	http.HandleFunc("/", home)

	// Passive Mobs
	for _, mob := range content.PassiveMobs {
		http.HandleFunc("/"+strings.ReplaceAll(mob, " ", "-")+"/", mobHandler(mob))
	}

	// Hostile Mobs
	for _, mob := range content.HostileMobs {
		http.HandleFunc("/"+strings.ReplaceAll(mob, " ", "-")+"/", mobHandler(mob))
	}

	log.Fatal(http.ListenAndServe(getPort(), nil))
}

// Render the template to the view
func renderTemplate(w http.ResponseWriter, p PageProps, check int) {
	if check == 0 {
		err := templates.ExecuteTemplate(w, "index.html", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if check == 1 {
		err := templates.ExecuteTemplate(w, "mob.html", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// Render Home Template
func home(w http.ResponseWriter, _ *http.Request) {
	p := PageProps{
		Title: "Minecraft Mobs",
	}
	renderTemplate(w, p, 0)
}

// Render Mob Template
func mobHandler(mob string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		c := content.GetMobContent(mob)
		p := PageProps{
			Title: 		"Mobs: " + strings.Title(mob),
			Mob:   		strings.Title(mob),
			Content:	c,
		}
		renderTemplate(w, p, 1)
	}
}
