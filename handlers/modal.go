package handlers

import (
	"net/http"

	"github.com/TylerGilman/nereus_main_site/views/components"
)

func HandleMoreOptionsModal(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	isHtmxRequest := r.Header.Get("HX-Request") == "true"
	return Render(w, r, components.Modal(components.MoreOptionsModalContent(), !isHtmxRequest))
}

func HandleNotificationsModal(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	isHtmxRequest := r.Header.Get("HX-Request") == "true"
	return Render(w, r, components.Modal(components.NotificationsModalContent(), !isHtmxRequest))
}

func HandleUserProfileModal(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	isHtmxRequest := r.Header.Get("HX-Request") == "true"
	return Render(w, r, components.Modal(components.UserProfileModalContent(), !isHtmxRequest))
}

func HandleCloseModal(w http.ResponseWriter, r *http.Request) error {
	isHtmxRequest := r.Header.Get("HX-Request") == "true"
	if isHtmxRequest {
		w.Header().Set("Content-Type", "text/html")
		_, err := w.Write([]byte("<div id=\"modal-container\"></div>"))
		return err
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}
}