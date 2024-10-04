package handler

import (
	"context"
	"database/sql"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"xerus/internal/db"
	"xerus/internal/template"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

type Handler struct {
	QueryStore *db.QueryStore
}

func NewHandler(dba *sql.DB) *Handler {
	return &Handler{
		QueryStore: db.NewQueryStore(dba),
	}
}

func (h *Handler) ServeStaticFiles(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path[len("/static/"):]
	fullPath := filepath.Join(".", "static", filePath)
	http.ServeFile(w, r, fullPath)
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	initialQueries := h.QueryStore.LoadQueries()
	template.Home("QueryFinder", initialQueries).Render(context.Background(), w)
}

func (h *Handler) AddQuery(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		template.ErrorToast(err.Error()).Render(r.Context(), w)
		return
	}

	vals, err := url.ParseQuery(string(body))
	if err != nil {
		template.ErrorToast(err.Error()).Render(r.Context(), w)
		return
	}

	if err := h.QueryStore.AddQuery(strings.ToLower(vals["queryName"][0]), strings.ToLower(vals["query"][0])); err != nil {
		template.ErrorToast(err.Error()).Render(r.Context(), w)
		return
	}
	w.Header().Add("HX-Trigger", "newQuery")
	w.WriteHeader(204)
}

func (h *Handler) DeleteQuery(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		template.ErrorToast(err.Error()).Render(r.Context(), w)
		return
	}

	vals, err := url.ParseQuery(string(body))
	if err != nil {
		template.ErrorToast(err.Error()).Render(r.Context(), w)
		return
	}

	h.QueryStore.DeleteQuery(vals["queryName"][0])
	w.Header().Add("HX-Trigger", "newDeletion")
	w.WriteHeader(204)
}

func (h *Handler) SearchQuery(w http.ResponseWriter, r *http.Request) {
	searchInput := r.URL.Query()["search"][0]
	ranks := fuzzy.RankFind(searchInput, h.QueryStore.GetQueries())
	var targets []string
	for _, item := range ranks {
		targets = append(targets, item.Target)
	}
	template.SearchResult(targets).Render(r.Context(), w)
}

func (h *Handler) GetQuery(w http.ResponseWriter, r *http.Request) {
	query, err := h.QueryStore.GetQuery(r.URL.Query()["queryName"][0])
	if err != nil {
		template.ErrorToast(err.Error()).Render(r.Context(), w)
		return
	}
	template.FullQueryView(query).Render(r.Context(), w)
}
