package handlers

import (
	"encoding/json"
	"go-gorm-orm/pkg/mocks"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Read dynamic ids
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Iterate over all the mock Books
	for index, book := range mocks.Books {
		if book.Id == id {
			// Delete book if id compare dynamic id
			mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted!")
			break
		}
	}
}
