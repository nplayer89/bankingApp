package domain

import "banking/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
