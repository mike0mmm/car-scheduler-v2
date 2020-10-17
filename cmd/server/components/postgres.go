package components

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/lib/pq"
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

// TODO: check if exists before inserting
func (p *postgres) SaveContactPerson(stopID, name, phone string) error {
	return nil
}

// TODO: check if exists before inserting
func (p *postgres) SaveCompany(company models.Company) error {
	query := `insert into company(id, address, city, phone, phone_secondary, email, name, logo_path) 
		values ($1, $2, $3, $4, $5, $6, $7, $8)`
	companyUUID := uuid.NewV4()
	_, err := p.db.Exec(query,
		companyUUID,
		company.Address,
		company.City,
		company.Phone,
		company.SeconadaryPhone,
		company.Email,
		company.CompanyName,
		company.LogoPath,
	)
	return err
}

func (p *postgres) GetCompany(name string) (models.Company, error) {
	c := models.Company{}
	query := fmt.Sprintf("select * from company where name='%s'", name)
	res := p.db.QueryRow(query)
	
	err := res.Scan(&c.Address, &c.City, &c.Phone, &c.SeconadaryPhone, &c.Email, &c.CompanyName, &c.LogoPath)
	return c, err
}

// TODO: check if exists before inserting
func (p *postgres) SaveUser(user models.User) error {
	query := `insert into "user"(username, user_id, role, profile_picture, phone_secondary, phone, 
		password, name, license_types, license_number, licence_expiration_date, family_name, expiration_date,
		email, date_of_birth, city, address, access_token, company) 
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)`

	res, err := p.db.Exec(query,
		user.Username,
		user.UserID,
		user.Role,
		user.ProfilePicture,
		user.SecondaryPhonee,
		user.Phone,
		user.Password,
		user.Name,
		pq.Array(user.LicenseTypes),
		user.LicenseNumber,
		user.LicenceExpirationDate,
		user.FamilyName,
		user.ExpirationDate,
		user.Email,
		user.DateOfBirth,
		user.City,
		user.Address,
		user.AccessToken,
		user.Company,
	)
	fmt.Println(res)
	return err
}

func (p *postgres) SaveCar(car models.Car) error {
	query := `insert into car(name, license_plate, model, manufacturing_year, description, 
		car_type, vehicle_license_expiration, insurance_expiration, last_treatment, last_brakes_check, capacity) 
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := p.db.Exec(query,
		car.Name,
		car.LicensePlate,
		car.Model,
		car.ManufacturingYear,
		car.Description,
		car.CarType,
		car.VehicleLicenseExpiration,
		car.InsuranceExpiration,
		car.LastTreatment,
		car.LastBrakesCheck,
		car.Capacity)
	return err
}

func (p *postgres) GetCar(licensePlate string) (models.Car, error) {
	query := fmt.Sprintf("select * from car where license_plate='%s'", licensePlate)

	car := models.Car{}
	row := p.db.QueryRow(query)
	err := row.Scan(&car.Name, &car.LicensePlate, &car.Model, &car.ManufacturingYear, &car.Description,
		 &car.CarType, &car.VehicleLicenseExpiration, &car.InsuranceExpiration, &car.LastTreatment, 
		 &car.LastBrakesCheck, &car.Capacity)
	
		 return car, err
}
