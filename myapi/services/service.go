package services

import "database/sql"

type MyAppService struct {
	db *sql.DB
}

// This is called Constructor
func NewMyAppService(db *sql.DB) *MyAppService {
	return &MyAppService{db: db}
}
