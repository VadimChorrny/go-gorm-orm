package handlers

import (
	"encoding/json"
	"go-gorm-orm/pkg/mocks"
	"go-gorm-orm/pkg/models"
	"io/ioutil"
	"log"
	mathRand "math/rand"
	"net/http"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	// Append to the Book mocks
	book.Id = mathRand.Intn(100)
	mocks.Books = append(mocks.Books, book)
	// Send a 201 created response
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode("Created!")

}
