package main

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

type Person struct {
	ID        int    `db:"user_id" json:"id,omitempty"`
	Username  string `db:"username" json:"username,omitempty"`
	Firstname string `db:"first_name" json:"first_name,omitempty"`
	Lastname  string `db:"last_name" json:"last_name,omitempty"`
	Gender    string `db:"gender" json:"gender,omitempty"`
	Password  string `db:"password" json:"password,omitempty"`
	Status    int    `db:"status" json:"status,omitempty"`
}

var people []Person

func main() {
	dsn := "root:@/sqlquery-poc"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	checkErr(err)

	r := mux.NewRouter()
	r.HandleFunc("/query/{query}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		query := vars["query"]

		w.Header().Set("Content-Type", "application/json")

		people := []Person{}
		err = db.Select(&people, query)
		if err != nil {
			err = json.NewEncoder(w).Encode(err)
		} else  {
			err = json.NewEncoder(w).Encode(people)
			checkErr(err)
		}
	})
	println("server listening on localhost:8000")
	err = http.ListenAndServe(":8000", r)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
