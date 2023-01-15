package handlers

import (
	"encoding/json"
	"go-gorm-orm/pkg/mocks"
	"go-gorm-orm/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Read dynamic ids
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var updatedBook models.Book
	json.Unmarshal(body, &updatedBook)
	// Iterate over all the mock Books
	for index, book := range mocks.Books {
		if book.Id == id {
			// Update book if match id equal dynamic id

			book.Title = updatedBook.Title
			book.Desc = updatedBook.Desc

			mocks.Books[index] = book

			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Updated!")
		}
	}
}
