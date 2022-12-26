package routes

import (
	"github.com/gorilla/mux"
	"github.com/sagarhande/go-bookstore/pkg/controllers"
)

var registerBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{bookID}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{bookID}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookID}", controllers.DeleteBook).Methods("DELETE")
}
