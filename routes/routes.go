package routes

import (
	"net/http"

	"rest-api-crud-books/controllers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the Book API!"))
	}).Methods("GET")

	// Book routes
	router.HandleFunc("/books", controllers.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBookByID(db)).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook(db)).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook(db)).Methods("DELETE")

	return router
}
