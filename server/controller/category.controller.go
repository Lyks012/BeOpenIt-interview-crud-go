package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"crud/models"

	"crud/config"
) 

type ResponseCategory struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data []model.Category `json:data`
}

func GetAllCategories(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	
	var category model.Category
	var arrCategory []model.Category
	var response ResponseCategory

	db := config.DBConnect()

	defer db.Close()

	rows, err := db.Query("SELECT * FROM category")

	if err != nil{
		log.Print(err)
	}

	for rows.Next(){
		err = rows.Scan(&category.Id, &category.Name)
		if err != nil{
			log.Fatal(err.Error())
		}else{
			arrCategory = append(arrCategory, category)
		}
	}

	response.Status = 200;
	response.Message = "succes"
	response.Data = arrCategory

    json.NewEncoder(w).Encode(response )
}

func CreateCategory(w http.ResponseWriter, r *http.Request){
	var res Response
	var category model.Category

	db := config.DBConnect()

	defer db.Close()

	
	invalidBodyError := json.NewDecoder(r.Body).Decode(&category)
	if invalidBodyError != nil {
        http.Error(w, invalidBodyError.Error(), http.StatusBadRequest)
        return
    }
	name := category.Name

	_, err := db.Exec("INSERT INTO category VALUES(?, ?)", "", name)

	if err != nil{
		log.Print(err)
		return
	}

	res.Status = 200
	res.Message = "Inserted with succes"
	

	w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Headers:", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
    json.NewEncoder(w).Encode(res)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request){
	var response Response

	db := config.DBConnect()

	defer db.Close()

	id := r.FormValue("id");

	_, err := db.Exec("DELETE FROM category WHERE id = ?", id)

	if err != nil{
		log.Println(err)
	}

	response.Status = 200
	response.Message = "Post deleted with succes"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}