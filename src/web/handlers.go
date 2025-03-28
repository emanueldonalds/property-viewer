package web

import (
    "strconv"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/emanueldonalds/property-viewer/db"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, sqldb *sql.DB) {
	listings := db.GetListings(r, sqldb)
	lastScrape := db.GetLastScrape(sqldb)
	index := Index(listings, lastScrape)
	index.Render(r.Context(), w)
}

func FilterHandler(w http.ResponseWriter, r *http.Request, sqldb *sql.DB) {
	listings := db.GetListings(r, sqldb)
	lastScrape := db.GetLastScrape(sqldb)
	index := Listings(listings, lastScrape)
	index.Render(r.Context(), w)
}

func DetailsHandler(sqldb *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
        idParam := params["id"]

        id, err := strconv.Atoi(idParam)

        if err != nil {
            fmt.Printf("Could not convert ID param [%s] to int", idParam)
        }

        listing := db.GetListing(id, sqldb);

        listingPage := Listing(listing)
        listingPage.Render(r.Context(), w)
	}
}
