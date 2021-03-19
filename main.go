package main

import (
	"html/template"
	"log"
	"net/http"
  "os"
)

type Page struct {
  Title string
  Body []byte
}

var templates = template.Must(template.ParseFiles("templates/index.html","templates/creeper.html", "templates/pig.html"))

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

  http.HandleFunc("/", Home)
  http.HandleFunc("/Pig/", Pig)
  http.HandleFunc("/Creeper/",Creeper)
  log.Fatal(http.ListenAndServe(getPort(), nil))
}

func renderTemplate(w http.ResponseWriter, tmpl string, p Page) {
  err := templates.ExecuteTemplate(w, tmpl+".html", p)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func Home(w http.ResponseWriter, r *http.Request) {
  p := Page{
    Title: "Supa Fast Fingers",
  }
  renderTemplate(w, "index", p)
}

func Pig(w http.ResponseWriter, r *http.Request)  {
  p := Page{
    Title: "Pig",
  }
  renderTemplate(w, "pig", p)
}

func Creeper(w http.ResponseWriter, r *http.Request)  {
  p := Page{
    Title: "Creeper",
  }
  renderTemplate(w, "creeper", p)
}
