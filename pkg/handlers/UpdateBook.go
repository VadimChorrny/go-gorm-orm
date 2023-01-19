package handlers

import (
	"encoding/json"
	"fmt"
	"go-gorm-orm/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
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

	var book models.Book

	if result := h.DB.Find(&book, id); result.Error != nil {
		fmt.Println(err)
	}

	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	book.Desc = updatedBook.Desc

	h.DB.Save(&book)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Updated!")
}
