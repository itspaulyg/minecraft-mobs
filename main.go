package main

import (
	"html/template"
	"log"
	"net/http"
  "os"
)

// Generic to each html page
type Page struct {
  Title string
	Mob		string
}

// Keep track of html files
var templates = template.Must(template.ParseGlob("templates/*.html"))
var passive_templates = template.Must(template.ParseGlob("templates/passive/*.html"))
var hostile_templates = template.Must(template.ParseGlob("templates/hostile/*.html"))

// Get port either local or heroku hosted
func getPort() string {
  p := os.Getenv("PORT")
  if p != "" {
    return ":" + p
  }
  return ":8080"
}

func main() {
  http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
  http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	// Home
  http.HandleFunc("/", Home)

	// Passive Mobs
	http.HandleFunc("/Chicken/", Chicken)
	http.HandleFunc("/Cow/", Cow)
  http.HandleFunc("/Pig/", Pig)
	http.HandleFunc("/Sheep/", Sheep)
	http.HandleFunc("/Villager/", Villager)

	// Hostile Mobs
	http.HandleFunc("/Blaze/",Blaze)
  http.HandleFunc("/Creeper/",Creeper)
	http.HandleFunc("/Ghast/",Ghast)
	http.HandleFunc("/MagmaCube/",MagmaCube)
	http.HandleFunc("/Skeleton/",Skeleton)
	http.HandleFunc("/Slime/",Slime)
	http.HandleFunc("/Zombie/",Zombie)

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

// Render Home Template

func Home(w http.ResponseWriter, r *http.Request) {
  p := Page{
    Title: "Minecraft Mobs",
  }
  renderTemplate(w, "index", p, 0)
}

// Passive Mob Templates

func Chicken(w http.ResponseWriter, r *http.Request)  {
  p := Page{
		Title: "Mobs: Chicken",
    Mob: "Chicken",
  }
  renderTemplate(w, "chicken", p, 1)
}

func Cow(w http.ResponseWriter, r *http.Request)  {
  p := Page{
		Title: "Mobs: Cow",
    Mob: "Cow",
  }
  renderTemplate(w, "cow", p, 1)
}

func Pig(w http.ResponseWriter, r *http.Request)  {
  p := Page{
		Title: "Mobs: Pig",
    Mob: "Pig",
  }
  renderTemplate(w, "pig", p, 1)
}

func Sheep(w http.ResponseWriter, r *http.Request)  {
  p := Page{
		Title: "Mobs: Sheep",
    Mob: "Sheep",
  }
  renderTemplate(w, "sheep", p, 1)
}

func Villager(w http.ResponseWriter, r *http.Request)  {
  p := Page{
		Title: "Mobs: Villager",
    Mob: "Villager",
  }
  renderTemplate(w, "villager", p, 1)
}

// Hostile Mob Templates

func Blaze(w http.ResponseWriter, r *http.Request)  {
  p := Page{
		Title: "Mobs: Blaze",
    Mob: "Blaze",
  }
  renderTemplate(w, "blaze", p, 2)
}

func Creeper(w http.ResponseWriter, r *http.Request)  {
  p := Page{
		Title: "Mobs: Creeper",
    Mob: "Creeper",
  }
  renderTemplate(w, "creeper", p, 2)
}

func Ghast(w http.ResponseWriter, r *http.Request)  {
  p := Page{
		Title: "Mobs: Ghast",
    Mob: "Ghast",
  }
  renderTemplate(w, "ghast", p, 2)
}

func MagmaCube(w http.ResponseWriter, r *http.Request)  {
  p := Page{
		Title: "Mobs: Magma Cube",
    Mob: "Magma Cube",
  }
  renderTemplate(w, "magmacube", p, 2)
}

func Skeleton(w http.ResponseWriter, r *http.Request)  {
  p := Page{
		Title: "Mobs: Skeleton",
    Mob: "Skeleton",
  }
  renderTemplate(w, "skeleton", p, 2)
}

func Slime(w http.ResponseWriter, r *http.Request)  {
  p := Page{
		Title: "Mobs: Slime",
    Mob: "Slime",
  }
  renderTemplate(w, "slime", p, 2)
}

func Zombie(w http.ResponseWriter, r *http.Request)  {
  p := Page{
		Title: "Mobs: Zombie",
    Mob: "Zombie",
  }
  renderTemplate(w, "zombie", p, 2)
}
