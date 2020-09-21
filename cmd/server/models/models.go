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

