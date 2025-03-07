package main

import (
	"fmt"
	"github.com/emanueldonalds/property-viewer/db"
	"github.com/emanueldonalds/property-viewer/rss"
	"github.com/emanueldonalds/property-viewer/web"
	"log"
	"net/http"
	"os"
)

func main() {
	assetsDir := "./assets"

	info, err := os.Stat(assetsDir)

	if err != nil {
		panic("Could not stat assets directory. Make sure assets dir is in the working directory.")
	}
	if info.Mode().Perm()&0444 != 0444 {
		panic("Missing permissions to read assets")
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(assetsDir))
	db := db.GetDb()

	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { web.IndexHandler(w, r, db) })
	mux.HandleFunc("/filter", func(w http.ResponseWriter, r *http.Request) { web.FilterHandler(w, r, db) })
	mux.HandleFunc("/rss", func(w http.ResponseWriter, r *http.Request) { rss.RssHandler(w, r, db) })

	fmt.Println("Listening on :4932")
	log.Fatal(http.ListenAndServe(":4932", mux))
}
