package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"fmt"
	
	"github.com/itspaulyg/minecraft-mobs/content"
	"github.com/itspaulyg/minecraft-mobs/model"
)

// Generic to each html page
type PageProps struct {
	Title   		string
	Mobs			[]string
	Mob     		string
	Content 		model.Content
	Year			string
	MobTypeFilter	MobTypeFilterOptions
}

type MobTypeFilterOptions struct {
	Options		[]string
	Selected	string
}

var mobTypeFilterOptions = MobTypeFilterOptions{
	Options:	[]string{"all", "passive", "neutral", "hostile"},
	Selected: 	"all",
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
	"findImgSize": func(height float32) int {
		return int(height * 100)
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

	// Filter
	http.HandleFunc("/filter", filter)

	// Mobs
	for _, mob := range content.AllMobs {
		http.HandleFunc("/"+strings.ReplaceAll(mob, " ", "-")+"/", mobHandler(mob))
	}

	log.Fatal(http.ListenAndServe(getPort(), nil))
}

// Render the template to the view
func renderTemplate(w http.ResponseWriter, p PageProps, template int) {
	if template == 0 {
		err := templates.ExecuteTemplate(w, "index.html", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if template == 1 {
		err := templates.ExecuteTemplate(w, "mob.html", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// Render Home Template
func home(w http.ResponseWriter, _ *http.Request) {
	currentTime := time.Now()
	currentYear := currentTime.Format("2006")

	mobTypeFilterOptions.Selected = "all"
	p := PageProps{
		Title:			"Minecraft Mobs",
		Mobs:			content.AllMobs,
		Year: 			currentYear,
		MobTypeFilter: 	mobTypeFilterOptions,
	}
	renderTemplate(w, p, 0)
}

func filter(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	currentYear := currentTime.Format("2006")

	r.ParseForm()

	fmt.Println(r.Form)

	mobs := content.GetMobsByFilter(r.Form["mobType"][0])
	mobTypeFilterOptions.Selected = r.Form["mobType"][0]

	p := PageProps{
		Title:			"Minecraft Mobs",
		Mobs:			mobs,
		Year: 			currentYear,
		MobTypeFilter: 	mobTypeFilterOptions,
	}

	renderTemplate(w, p, 0)
}

// Render Mob Template
func mobHandler(mob string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		currentTime := time.Now()
		currentYear := currentTime.Format("2006")
		
		c := content.GetMobContent(mob)
		p := PageProps{
			Title: 			"Mobs: " + strings.Title(mob),
			Mob:   			strings.Title(mob),
			Content:		c,
			Year:			currentYear,
		}
		renderTemplate(w, p, 1)
	}
}
