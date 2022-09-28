package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"crud/config"
	model "crud/models"

	"github.com/gorilla/mux"
)

type Response struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    []model.Post `json:data`
}

func GetPostsByCategoryId(w http.ResponseWriter, r *http.Request) {
	var post model.Post
	var arrPost []model.Post

	db := config.DBConnect()

	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]
	// id := r.URL.Query().Get("id")
	log.Print("id" + id)
	fmt.Println(id)
	rows, err := db.Query("SELECT * FROM post WHERE id_category = ?", id)

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&post.Id, &post.Title, &post.Description, &post.Category_id)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrPost = append(arrPost, post)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(arrPost)

}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	var post model.Post
	var arrPost []model.Post
	var response Response

	db := config.DBConnect()

	defer db.Close()

	rows, err := db.Query("SELECT post.id, post.title, post.description, post.id_category, category.name FROM post INNER JOIN category ON post.id_category = category.id")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&post.Id, &post.Title, &post.Description, &post.Category_id, &post.Category_name)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrPost = append(arrPost, post)
		}
	}

	response.Status = 200
	response.Message = "succes"
	response.Data = arrPost

	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	var post model.Post
	var arrPost []model.Post
	var response Response

	db := config.DBConnect()

	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	rows, err := db.Query("SELECT * FROM post WHERE id = ?", id)

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&post.Id, &post.Title, &post.Description, &post.Category_id)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrPost = append(arrPost, post)
		}
	}

	response.Status = 200
	response.Message = "succes"
	response.Data = arrPost

	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var res Response
	var post model.Post

	invalidBodyError := json.NewDecoder(r.Body).Decode(&post)
	if invalidBodyError != nil {
		http.Error(w, invalidBodyError.Error(), http.StatusBadRequest)
		return
	}

	db := config.DBConnect()

	defer db.Close()

	title := post.Title
	description := post.Description
	categoryId := post.Category_id

	fmt.Println(categoryId)

	_, err := db.Exec("INSERT INTO post VALUES(?, ?, ?, ?)", 0, title, description, categoryId)

	if err != nil {
		log.Print(err)
		return
	}

	res.Status = 200
	res.Message = "Inserted with succes"

	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(res)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting")
	var response Response

	db := config.DBConnect()

	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	_, err := db.Exec("DELETE FROM post WHERE id = ?", id)

	if err != nil {
		log.Println(err)
	}

	response.Status = 200
	response.Message = "Post deleted with succes"

	w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
