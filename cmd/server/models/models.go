package models

type Company struct {
    Address string
    City string
    Email string
    LogoPath string `json:"logo_path"`
    CompanyName string `json:"company_name"`
    Phone string
    SeconadaryPhone string `json:"phone_secondary"`
}

type User struct {
    Username string
    UserId string `json:"user_id"`
    Role string
    SecondaryPhonee string `json:"phone_secondary"`
    Phone string 
    Password string
    Name string
    LicenseTypes string `json:"license_types"`
    LicenseNumber string `json:"license_number"`
    LicenceExpirationDate string `json:"licence_expiration_date"`
    FamilyName string `json:"family_name"`
    ExpirationDate string `json:"expiration_date"`
    Email string
    DateOfBirth string `json:"date_of_birth"`
    City string
    Address string
    AccessToken string `json:"access_token"`
}