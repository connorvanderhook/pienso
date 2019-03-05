package middleware

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	gcs "github.com/connorvanderhook/pienso/api"
)

func GetPost(w http.ResponseWriter, r *http.Request) {
	var err error
	var article []byte

	if articleID := chi.URLParam(r, "articleID"); articleID != "" {
		article, err = gcs.ReadFromBucket(toFileName(articleID))
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(422), 422)
			return
		}
	}
	w.Write(article)
}

func PostsByDate(w http.ResponseWriter, r *http.Request) {
	// Stub to get posts by date
}

func LastFivePosts(w http.ResponseWriter, r *http.Request) {
	// Stub for default Load of All Posts when blog is loaded
}

func toFileName(id string) string {
	return fmt.Sprintf("%s.md", id)
}
