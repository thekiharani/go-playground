package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title   string `json:"Title"`
	Summary string `json:"Summary"`
	Body    string `json:"Body"`
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts []Post
	DB.Find(&posts)
	json.NewEncoder(w).Encode(posts)
}

func ViewPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var post Post
	DB.First(&post, params["id"])
	json.NewEncoder(w).Encode(post)
}
func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	json.NewDecoder(r.Body).Decode(&post)
	DB.Create(&post)
	json.NewEncoder(w).Encode(post)
}
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var post Post
	DB.First(&post, params["id"])
	json.NewDecoder(r.Body).Decode(&post)
	DB.Save(&post)
	json.NewEncoder(w).Encode(post)
}
func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var post Post
	DB.Delete(&post, params["id"])
	json.NewEncoder(w).Encode("Deleted")
}
