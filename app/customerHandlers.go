package app

import (
	"banking/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers, err := ch.service.GetAllCustomers()
	if err != nil {
		writeResponseJson(w, err.Code, err.AsMessage())
	} else {
		writeResponseJson(w, http.StatusOK, customers)
	}

}

func (ch *CustomerHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponseJson(w, err.Code, err.AsMessage())
	} else {
		writeResponseJson(w, http.StatusOK, customer)
	}
}

func writeResponseJson(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Conent-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
