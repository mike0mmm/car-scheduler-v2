package models

import "time"

type Company struct {
    Address string
    City string
    Email string
    LogoPath string `json:"logo_path" db:"logo_path"`
    CompanyName string `json:"company_name" db:"name"`
    Phone string
    SeconadaryPhone string `json:"phone_secondary" db:"phone_secondary"`
}

type User struct {
    Username string
    UserID int64 `json:"user_id" db:"user_id"`
    Role string
    DateOfBirth time.Time `json:"date_of_birth" db:"date_of_birth"`
    ProfilePicture string `josn:"profile_picture" db:"profile_picture"`
    SecondaryPhonee string `json:"phone_secondary" db:"phone_secondary"`
    Phone string 
    Password string
    Name string
    LicenseTypes []string `json:"license_types" db:"license_types"`
    LicenseNumber string `json:"license_number" db:"license_number"`
    LicenceExpirationDate time.Time `json:"licence_expiration_date" db:"licence_expiration_date"`
    FamilyName string `json:"family_name" db:"family_name"`
    ExpirationDate time.Time `json:"expiration_date" db:"expiration_date"`
    Email string
    City string
    Address string
    AccessToken string `json:"access_token" db:"access_token"`
    Company string
}

type Car struct {
    CarID string `json:"car_id" db:"car_id"`
    Name string
    LicensePlate string `json:"license_plate" db:"license_plate"`
    Model string
    ManufacturingYear string `json:manufacturing_year" db:manufacturing_year"`
    Description string
    CarType string `json:"car_type" db:"car_type"`
    VehicleLicenseExpiration string `json:"vehicle_license_expiration" db:"vehicle_license_expiration"`
    InsuranceExpiration string `json:"insurance_expiration" db:"insurance_expiration"`
    LastTreatment string `json:"last_treatment" db:"last_treatment"`
    LastBrakesCheck string `json:"last_brakes_check" db:"last_brakes_check"`
    Capacity string
}








type ContactPerson struct {
    Name string
    Phone string
}

type Ride struct {
    ID string
    Type string
    Price string
    RideDate string `json:"ride_date" db:"ride_date"`
    StartTime string `json:"start_time" db:"start_time"`
    EndTime string `json:"end_time" db:"end_time"`
    Reaccuring string
    Stops string
}

type ScheduleEntry struct {
    StartSime string `json:"start_time" db:"start_time"`
    EndTime string `json:"end_time" db:"end_time"`
    RideDate string `json:"ride_date" db:"ride_date"`
    Driver string
    Car string
    Accepted string
    RideDetails string `json:"ride_details" db:"ride_details"`
}

type Stop struct {
    ID string
    Name string
    Address string
    NumberOfPassangers string `json:"number_of_passangers" db:"number_of_passangers"`
    StopTime string `json:"stop_time" db:"stop_time"`
    Note string
    Contact string
}