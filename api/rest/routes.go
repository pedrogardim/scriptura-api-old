package routes

import (
	"net/http"

	"github.com/pedrogardim/scriptura-api/api/rest/handlers"
)

func Init() {
	http.HandleFunc("/api/verse/", handlers.GetVerse)

}
