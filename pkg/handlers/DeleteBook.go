package handlers

import (
	"encoding/json"
	"fmt"
	"go-gorm-orm/pkg/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Read dynamic ids
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find the book by Id
	var book models.Book
	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Delete book
	h.DB.Delete(&book)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted!")
}
