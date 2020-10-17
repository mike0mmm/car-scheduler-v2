package persister

import "github.com/mike0mmm/car-scheduler-v2/cmd/server/models"

type Persister interface {
	SaveContactPerson(stopID, name, phone string) error
	
	SaveCompany(models.Company) error
	GetCompany(string) (models.Company, error)

	SaveUser(models.User) error
	GetUser(string) (models.User, error)

	SaveCar(models.Car) error
	GetCar(string) (models.Car, error)
}
