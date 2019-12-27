//go:generate go run views/generate.go

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/majomi/soapberry/static"
	"github.com/majomi/soapberry/views"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.DefaultCompress)

	srv := &http.Server{
		Addr:         ":3000",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	homeView := views.NewView("base", "home")
	contactView := views.NewView("base", "contact")
	blubView := views.NewView("base", "blub")

	fs := static.FileSystem{
		Fs: static.Assets,
	}

	r.Handle("/static/*", http.FileServer(fs))

	r.Get("/blub", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		err := blubView.Template.ExecuteTemplate(w, blubView.Layout, nil)
		if err != nil {
			panic(err)
		}
	})

	r.Get("/contact", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
		if err != nil {
			panic(err)
		}
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
		if err != nil {
			panic(err)
		}
	})

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
