package pladelivery

import (
	"codex/internal/pkg/domain"

	"codex/internal/pkg/utils/sanitizer"
	"encoding/json"
	"net/http"
)

func (handler *PlaylistHandler) CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	playlistRequest := new(domain.PlaylistRequest)
	err := json.NewDecoder(r.Body).Decode(&playlistRequest)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	sanitizer.SanitizePlaylistCreating(playlistRequest)

	us, err := handler.PlaylistUsecase.CreatePlaylist(*playlistRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(us)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(out)
}

func (handler *PlaylistHandler) AddMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	addPlaylistInfo := new(domain.MovieInPlaylist)
	err := json.NewDecoder(r.Body).Decode(&addPlaylistInfo)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	err = handler.PlaylistUsecase.AddMovie(*addPlaylistInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (handler *PlaylistHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	MovieInPlaylist := new(domain.MovieInPlaylist)
	err := json.NewDecoder(r.Body).Decode(&MovieInPlaylist)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	err = handler.PlaylistUsecase.DeleteMovie(*MovieInPlaylist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *PlaylistHandler) DeletePlaylist(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	deletePlaylistInfo := new(domain.DeletePlaylistInfo)
	err := json.NewDecoder(r.Body).Decode(&deletePlaylistInfo)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	err = handler.PlaylistUsecase.DeletePlaylist(*deletePlaylistInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
