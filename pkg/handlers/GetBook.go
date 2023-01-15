package handlers

import (
	"encoding/json"
	"go-gorm-orm/pkg/mocks"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	// Iterate over all the mock book
	for _, book := range mocks.Books {
		if book.Id == id {
			// If ids are equal send book as response
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			break
		}
	}
}
