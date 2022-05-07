package main

import (
	"context"
	// "database/sql"
	"encoding/json"
	// "fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

var ctx = context.Background()

type Connection struct {
	db *sqlx.DB
}

type Person struct {
	FirstName  string `json:"name" db:"first_name"`
	SecondName string `json:"secondname" db:"second_name"`
	Age        string `json:"age" db:"age"`
	ID         string `json:"id" db:"id"`
}

func (s *Connection) getAllPersons(w http.ResponseWriter, r *http.Request) {
	var persons []Person
	w.Header().Set("Content-Type", "application/json")
	query := "SELECT * FROM persons;"
	err := s.db.Select(&persons, query)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(persons)
}

func (s *Connection) getPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var persons []Person
	query := "SELECT * FROM persons;"
	err := s.db.Select(&persons, query)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range persons {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func (s *Connection) createPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person Person
	person.ID = strconv.Itoa(rand.Intn(100000))
	query := `INSERT INTO persons VALUES ($1, $2, $3, $4);`
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	_ = json.NewDecoder(r.Body).Decode(&person)
	rows, err := s.db.QueryContext(ctx, query, person.FirstName, person.SecondName, person.Age, person.ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	json.NewEncoder(w).Encode(person)
}

func (s *Connection) deletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	query := `DELETE FROM persons WHERE id = $1`
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	_, err := s.db.QueryContext(ctx, query, params["id"])
	if err != nil {
		log.Fatal(err)
	}
	var persons []Person
	query = "SELECT * FROM persons;"
	err = s.db.Select(&persons, query)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(persons)
}

func main() {

	db, openErr := sqlx.Connect("pgx", "postgres://user:mypassword@postgres:5432")
	if openErr != nil {
		log.Fatal("sql.Open(): %v", openErr)
	}
	s := Connection{}
	s.db = db
	_, err := db.Exec(`DROP TABLE IF EXISTS persons;`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`CREATE TABLE persons (
		first_name	varchar(80),
		second_name	varchar(80),
		age			varchar(80),
		id			varchar(80)
	
	);`)
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/persons", s.getAllPersons).Methods("GET")
	r.HandleFunc("/persons/{id}", s.getPerson).Methods("GET")
	r.HandleFunc("/persons", s.createPerson).Methods("POST")
	r.HandleFunc("/persons/{id}", s.deletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
