package domain

import "time"

// Policy represents a policy model
type Policy struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	StartDate      time.Time `json:"start_date"`
	DurationYears  int       `json:"duration_years"`
	Company        string    `json:"company"`
	InitialDeposit float64   `json:"initial_deposit"`
	Type           string    `json:"type"`
	UserTypes      []string  `json:"user_types"`
	TermsPerYear   int       `json:"terms_per_year"`
	TermAmount     float64   `json:"term_amount"`
	Interest       float64   `json:"interest"`
}
