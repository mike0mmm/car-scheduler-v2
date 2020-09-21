package persister

import "github.com/mike0mmm/car-scheduler-v2/cmd/server/models"

type Persister interface {
	SaveContactPerson(stopID, name, phone string) error
	SaveCompany(models.Company) error
}
