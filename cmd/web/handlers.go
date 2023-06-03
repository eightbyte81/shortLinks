package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"shortLinks/internal/model"
	"shortLinks/internal/service"
)

type HandlerFunc interface {
	Home(w http.ResponseWriter, r *http.Request)
	GetShortLink(w http.ResponseWriter, r *http.Request)
	GetDefaultLink(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	HandlerFunc
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles("./ui/html/pages/home.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
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
			http.Error(w, "Internal Server Error", 500)
		}
	}(r.Body)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	_, err = h.services.SetDefaultLink(model.Link{LinkData: string(body)})
	if err != nil {
		return
	}
	shortLinkData, err := h.services.GetShortLinkByDefaultLink(model.Link{LinkData: string(body)})

	_, err = w.Write([]byte(shortLinkData))
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (h *Handler) GetDefaultLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte(r.URL.Query().Get("sl")))
}
