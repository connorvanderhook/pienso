package main

import (
	"net/http"
	"time"

	sys "github.com/connorvanderhook/pienso/app/configuration"
	m "github.com/connorvanderhook/pienso/app/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Application holds global variables to be used in each request
type Application struct {
	Configuration *sys.Configuration
}

func newApplication(templatePath string) *Application {
	c := &sys.Configuration{
		PublicPath:    "public",
		TemplatePath:  templatePath,
		IsDevelopment: false,
	}
	return &Application{Configuration: c}
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(300 * time.Second))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	r.Route("/posts", func(r chi.Router) {
		r.Get("/{articleID}", m.GetPost)
		// r.With(paginate).Get("/", lastFivePosts)
		// r.Get("/{postSlug:[a-z-]+}", getPost)
		// r.With(paginate).Get("/{month}-{date}-{year}", postsByDate)
	})

	http.ListenAndServe(":3333", r)
}

// paginate is a stub to handle the request params for handling a paginated request.
func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
