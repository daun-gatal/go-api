package main // Import main

import (
	"go-api/database"
	"go-api/entity"
	"go-api/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Required for MySQL dialect
)

func main() {
	InitDB()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	InitialiseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))

}

func InitialiseHandlers(router *mux.Router) {
	router.HandleFunc("/create", handlers.CreatePerson).Methods("POST")
	router.HandleFunc("/get", handlers.GetAllPerson).Methods("GET")

}

func InitDB() {
	config := database.Config{
		ServerName: "localhost:3306",
		User:       "root",
		Password:   "root",
		DB:         "dev",
	}
	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Person{})
}
