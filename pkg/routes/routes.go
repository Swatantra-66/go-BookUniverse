package routes

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/Swatantra-66/go-bookstore/pkg/handlers"
	"github.com/gorilla/mux"
)

func getSafeDir(path string) string {
	dir, _ := os.Getwd()
	return filepath.Join(dir, path)
}

// book routing
func RegisterBookRoutes(router *mux.Router) {
	router.HandleFunc("/signup", handlers.SignUp).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")

	router.HandleFunc("/book", handlers.GetBookByUser).Methods("GET")
	router.HandleFunc("/book/{bookId}", handlers.GetBookById).Methods("GET")
	router.HandleFunc("/book", handlers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", handlers.DeleteBook).Methods("DELETE")

	router.HandleFunc("/u/{username}", handlers.ServePublicPage).Methods("GET")
	router.HandleFunc("/api/public/{username}", handlers.GetPublicBooks).Methods("GET")
	router.HandleFunc("/recommend", handlers.GetAIRecommendations).Methods("GET")
	router.HandleFunc("/api/magic-details", handlers.GetBookDetailsAI).Methods("GET")
	router.HandleFunc("/api/user/update", handlers.UpdateUser).Methods("POST")
	router.HandleFunc("/api/books/reset", handlers.ResetLibrary).Methods("DELETE")
	router.HandleFunc("/api/user/password", handlers.UpdatePassword).Methods("POST")

	cssDir := getSafeDir("css")
	jsDir := getSafeDir("js")

	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(cssDir))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir(jsDir))))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	staticDir := getSafeDir("static")
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(staticDir))))
}
