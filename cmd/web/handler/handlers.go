package handler

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"shortLinks/internal/model"
	"shortLinks/internal/service"
)

type Handlers interface {
	Home(w http.ResponseWriter, r *http.Request)
	HandleShortUrlRedirect(w http.ResponseWriter, r *http.Request)
	HandleShortUrlRedirectWithCache(w http.ResponseWriter, r *http.Request)
	GetShortLink(w http.ResponseWriter, r *http.Request)
	GetDefaultLink(w http.ResponseWriter, r *http.Request)
	GetShortLinkFromCache(w http.ResponseWriter, r *http.Request)
	GetDefaultLinkFromCache(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	Handlers
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/pages/home.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *Handler) HandleShortUrlRedirect(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "http://localhost:8080/home", http.StatusSeeOther)
		return
	}

	defaultLink, err := h.services.GetDefaultLinkByShortLinkData(r.Host + r.URL.Path)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	http.Redirect(w, r, defaultLink.LinkData, http.StatusSeeOther)
}

func (h *Handler) HandleShortUrlRedirectWithCache(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "http://localhost:8080/home", http.StatusSeeOther)
		return
	}

	defaultLink, err := h.services.GetDefaultLinkFromCacheByShortLink(r.Host + r.URL.Path)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	http.Redirect(w, r, defaultLink, http.StatusSeeOther)
}

func (h *Handler) GetShortLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}(r.Body)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	shortLink, err := h.services.GetShortLinkByDefaultLink(model.Link{LinkData: string(body)})
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	_, err = w.Write([]byte(shortLink.LinkData))
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *Handler) GetDefaultLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	defaultLink, err := h.services.GetDefaultLinkByShortLinkData(r.URL.Query().Get("sl"))
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	_, err = w.Write([]byte(defaultLink.LinkData))
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *Handler) GetShortLinkFromCache(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}(r.Body)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	shortLink, err := h.services.SetLinksInCache(model.Link{LinkData: string(body)})
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	_, err = w.Write([]byte(shortLink))
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *Handler) GetDefaultLinkFromCache(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	defaultLink, err := h.services.GetDefaultLinkFromCacheByShortLink(r.URL.Query().Get("sl"))
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	_, err = w.Write([]byte(defaultLink))
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
