package app

import (
	"banking/domain"
	"banking/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func Start() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Could not load .env")
	}
	checkEnv()

	router := mux.NewRouter()

	//wiring
	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositorySub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomer).Methods(http.MethodGet)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func checkEnv() {
	if os.Getenv("SERVER_ADDRESS") == "" {
		log.Fatal("Environment variable SERVER_ADDRESS is not defined")
	}
	if os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable SERVER_PORT is not defined")
	}
	if os.Getenv("DB_USER") == "" {
		log.Fatal("Environment variable DB_USER is not defined")
	}
	if os.Getenv("DB_PASSWD") == "" {
		log.Fatal("Environment variable DB_PASSWD is not defined")
	}
	if os.Getenv("DB_ADDR") == "" {
		log.Fatal("Environment variable DB_ADDR is not defined")
	}
	if os.Getenv("DB_PORT") == "" {
		log.Fatal("Environment variable DB_PORT is not defined")
	}
	if os.Getenv("DB_NAME") == "" {
		log.Fatal("Environment variable DB_NAME is not defined")
	}
}
