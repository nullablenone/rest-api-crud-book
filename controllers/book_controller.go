package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api-crud-books/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetBooks(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var books []models.Book

		w.Header().Set("Content-Type", "application/json")

		if err := db.Find(&books).Error; err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		response := map[string]interface{}{
			"status":  http.StatusOK,
			"message": "success fetching data",
			"data":    books,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func GetBookByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var book models.Book

		w.Header().Set("Content-Type", "application/json")

		if err := db.First(&book, params["id"]).Error; err != nil {
			w.WriteHeader(http.StatusNotFound)
			response := map[string]interface{}{
				"status":  http.StatusNotFound,
				"message": "Book not found",
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		response := map[string]interface{}{
			"status":  http.StatusOK,
			"message": "success fetching data",
			"data":    book,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func CreateBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		if err := db.Create(&book).Error; err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		response := map[string]interface{}{
			"status":  http.StatusCreated,
			"message": "success creating data",
			"data":    book,
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func UpdateBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var book models.Book

		w.Header().Set("Content-Type", "application/json")

		if err := db.First(&book, params["id"]).Error; err != nil {
			w.WriteHeader(http.StatusNotFound)
			response := map[string]interface{}{
				"status":  http.StatusNotFound,
				"message": "Book not found",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		if err := db.Save(&book).Error; err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		response := map[string]interface{}{
			"status":  http.StatusOK,
			"message": "success updating data",
			"data":    book,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func DeleteBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		w.Header().Set("Content-Type", "application/json")

		if err := db.Delete(&models.Book{}, params["id"]).Error; err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		response := map[string]interface{}{
			"status":  http.StatusNoContent,
			"message": "success deleting data",
		}
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(response)
	}
}
