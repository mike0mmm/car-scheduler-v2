package components

import (
	"database/sql"
	"fmt"
	"os"
	"github.com/mike0mmm/car-scheduler-v2/cmd/server/models"
	persister "github.com/mike0mmm/car-scheduler-v2/cmd/server/persiter"
	uuid "github.com/satori/go.uuid"
)

type postgres struct {
	db *sql.DB
}

func NewPersisterComponent() (persister.Persister, error) {
	psqlInfo := os.Getenv("DATABASE_URL")
	if len(psqlInfo) == 0 {
		psqlInfo = fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=require",
			os.Getenv("LOCAL_DATABASE_URL"),
			os.Getenv("LOCAL_DATABASE_PORT"),
			os.Getenv("LOCAL_DATABASE_USER"),
			os.Getenv("LOCAL_DATABASE_PASSWORD"),
			os.Getenv("LOCAL_DATABASE_NAME"))
	}

		
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to open connection. error: %s", err.Error()) 
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping. error: %s", err.Error())
	}

	return &postgres{db: db}, nil
}

func (p *postgres) SaveContactPerson(stopID, name, phone string) error {
	return nil
}

func (p *postgres) SaveCompany(company models.Company) error {

	query := `insert into company(id, address, city, phone, phone_secondary, email, name) 
		values ($1, $2, $3, $4, $5, $6, $7)`
	companyUUID := uuid.NewV4()
	_, err := p.db.Exec(query, 
		companyUUID, 
		company.Address,
		company.City,
		company.Phone,
		company.SeconadaryPhone,
		company.Email,
		company.CompanyName)
	return err
}