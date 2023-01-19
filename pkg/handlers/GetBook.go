package handlers

import (
	"encoding/json"
	"fmt"
	"go-gorm-orm/pkg/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) GetBook(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find book by Id
	var book models.Book
	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
