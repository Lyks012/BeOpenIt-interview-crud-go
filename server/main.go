package main

import (
	"fmt"
	"log"
	"net/http"

	// "database/sql"

	"github.com/gorilla/mux"

	"crud/controller"

	_ "github.com/go-sql-driver/mysql"
)

type WithCORS struct {
	r *mux.Router
}

func main() {

	router := mux.NewRouter()

	router.Use(CORS)

	router.HandleFunc("/posts", controller.GetAllPosts).Methods("GET", "OPTIONS")
	router.HandleFunc("/post/{id}", controller.GetPostById).Methods("GET", "OPTIONS")
	router.HandleFunc("/createPost", controller.CreatePost).Methods("POST", "OPTIONS")
	router.HandleFunc("/post/{id}", controller.DeletePost).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/category/{id}/posts", controller.GetPostsByCategoryId).Methods("GET", "OPTIONS")

	router.HandleFunc("/category", controller.GetAllCategories).Methods("GET", "OPTIONS")
	router.HandleFunc("/category", controller.CreateCategory).Methods("POST", "OPTIONS")
	http.Handle("/", router)

	fmt.Println("connected to port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))

}
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set headers
		w.Header().Set("Access-Control-Allow-Headers:", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		fmt.Println("okay")

		// Next
		next.ServeHTTP(w, r)
		return
	})
}
