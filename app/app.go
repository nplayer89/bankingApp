package app

import (
	"banking/domain"
	"banking/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	//wiring
	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositorySub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
