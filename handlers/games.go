package handlers

import (
	"net/http"

	"github.com/TylerGilman/nereus_main_site/views/games"
)

func HandleGames(w http.ResponseWriter, r *http.Request) error {
	isHtmxRequest := r.Header.Get("HX-Request") == "true"
	if isHtmxRequest {
		return Render(w, r, games.Partial())
	} else {
		return Render(w, r, games.Games())
	}
}