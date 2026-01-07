package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"groupie-tracker/models"
	"groupie-tracker/services"
)

type PageData struct {
	Query    string
	Searched bool
	Artists  []models.Artist
	NoResult bool
}

func filterArtists(artists []models.Artist, query string) []models.Artist {
	query = strings.ToLower(query)
	var results []models.Artist

	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), query) {
			results = append(results, artist)
		}
	}
	return results
}

func Home(w http.ResponseWriter, r *http.Request) {
	allArtists, err := services.GetArtists()
	if err != nil {
		log.Println(err)
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}

	data := PageData{Artists: allArtists}

	if r.Method == http.MethodPost {
		query := strings.TrimSpace(r.FormValue("group"))
		if query != "" {
			data.Query = query
			data.Searched = true
			data.Artists = filterArtists(allArtists, query)

			if len(data.Artists) == 0 {
			data.NoResult = true
		}
		}
	}

	tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "home.html")))
	tmpl.Execute(w, data)
}
