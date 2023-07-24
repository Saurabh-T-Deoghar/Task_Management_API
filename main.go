package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gorilla/mux"
)

var task []Task = []Task{}
var db *sql.DB
var err error

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"duedate"`
	Status      string `json:"status"`
}

func addItem(q http.ResponseWriter, r *http.Request) {

	stmt, err := db.Prepare("INSERT INTO Task(Title,Description,Due_Date) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	title := keyVal["title"]
	_, err = stmt.Exec(title)

	description := keyVal["Description"]
	_, err = stmt.Exec(description)

	due_date := keyVal["Due_Date"]
	_, err = stmt.Exec(due_date)

	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(q, "New post was created")

}

func getAllTasks(q http.ResponseWriter, r *http.Request) {
	q.Header().Set("Content-Type", "application/json")
	json.NewEncoder(q).Encode(task)

}

func getTasks(q http.ResponseWriter, r *http.Request) {

	q.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT id, title FROM posts WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var post Task
	for result.Next() {
		err := result.Scan(&post.ID, &post.Title)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(q).Encode(post)
}

func updateTasks(q http.ResponseWriter, r *http.Request) {

	q.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE posts SET title = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newTitle := keyVal["title"]
	_, err = stmt.Exec(newTitle, params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(q, "Post with ID = %s was updated", params["id"])
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Post with ID = %s was deleted", params["id"])
}

func main() {
	db, err = sql.Open("sqlite3", "mydb.db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/tasks", addItem).Methods("POST")
	router.HandleFunc("/tasks", getAllTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", getTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", updateTasks).Methods("PUT")

	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	http.ListenAndServe(":5000", router)

}
