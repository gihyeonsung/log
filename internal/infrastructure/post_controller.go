package infrastructure

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gihyeonsung/log/internal/application"
	"github.com/gihyeonsung/log/internal/domain"
)

type PostController struct {
	mux                *http.ServeMux
	postCreate         *application.PostCreate
	postDelete         *application.PostDelete
	postDocumentSearch *application.PostDocumentSearch
	postDocumentSync   *application.PostDocumentSync
	postFind           *application.PostFind
	postUpdate         *application.PostUpdate
}

func NewPostController(
	mux *http.ServeMux,
	postCreate *application.PostCreate,
	postDelete *application.PostDelete,
	postDocumentSearch *application.PostDocumentSearch,
	postDocumentSync *application.PostDocumentSync,
	postFind *application.PostFind,
	postUpdate *application.PostUpdate,
) *PostController {
	c := &PostController{
		mux:                mux,
		postCreate:         postCreate,
		postDelete:         postDelete,
		postDocumentSearch: postDocumentSearch,
		postDocumentSync:   postDocumentSync,
		postFind:           postFind,
		postUpdate:         postUpdate,
	}
	mux.Handle("/api/v1/posts", http.HandlerFunc(c.handle))
	return c
}

func (c *PostController) handle(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch {
	case path == "/api/v1/posts" && r.Method == http.MethodGet:
		c.handleGetPosts(w, r)
	case path == "/api/v1/posts" && r.Method == http.MethodPost:
		c.handlePostPosts(w, r)
	case len(path) > len("/api/v1/posts/") && r.Method == http.MethodPost && pathHasSuffix(path, "/update"):
		c.handlePostPostsIdUpdate(w, r)
	case len(path) > len("/api/v1/posts/") && r.Method == http.MethodDelete:
		c.handleDeletePostsId(w, r)
	case path == "/api/v1/posts/search" && r.Method == http.MethodGet:
		c.handleGetPostsSearch(w, r)
	default:
		http.Error(w, "not found", http.StatusNotFound)
	}
}

func pathHasSuffix(path, suffix string) bool {
	if len(path) < len(suffix) {
		return false
	}
	return path[len(path)-len(suffix):] == suffix
}

func (c *PostController) handleGetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := c.postFind.Exec()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if posts == nil {
		posts = []*domain.Post{}
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *PostController) handlePostPosts(w http.ResponseWriter, r *http.Request) {
	key := r.Header.Get("authorization")

	err := c.postCreate.Exec(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *PostController) handlePostPostsIdUpdate(w http.ResponseWriter, r *http.Request) {
	idStr := extractIDFromPath(r.URL.Path, "/posts/", "/update")
	if idStr == "" {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	postID, err := domain.PostIDFromString(idStr)
	if err != nil {
		http.Error(w, "invalid id format", http.StatusBadRequest)
		return
	}
	var updateReq struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	key := r.Header.Get("Authorization")
	if key == "" {
		http.Error(w, "missing authorization", http.StatusUnauthorized)
		return
	}
	err = c.postUpdate.Exec(key, postID, updateReq.Title, updateReq.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	go func() {
		err := c.postDocumentSync.Exec(postID)
		if err != nil {
			log.Printf("post document sync: %v", err)
		}
	}()

	w.WriteHeader(http.StatusOK)
}

func (c *PostController) handleDeletePostsId(w http.ResponseWriter, r *http.Request) {
	idStr := extractIDFromPath(r.URL.Path, "/posts/", "")
	if idStr == "" {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	postID, err := domain.PostIDFromString(idStr)
	if err != nil {
		http.Error(w, "invalid id format", http.StatusBadRequest)
		return
	}
	key := r.Header.Get("Authorization")
	if key == "" {
		http.Error(w, "missing authorization", http.StatusUnauthorized)
		return
	}
	err = c.postDelete.Exec(key, postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *PostController) handleGetPostsSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	posts, err := c.postDocumentSearch.Exec(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if posts == nil {
		posts = []*domain.PostDocument{}
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func extractIDFromPath(path, prefix, suffix string) string {
	start := len(prefix)
	end := len(path)
	if suffix != "" && pathHasSuffix(path, suffix) {
		end -= len(suffix)
	}
	if start >= end {
		return ""
	}
	return path[start:end]
}
