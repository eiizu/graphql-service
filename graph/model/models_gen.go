// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Appointment struct {
	ID       string    `json:"id"`
	Date     string    `json:"date"`
	Patient  *Patient  `json:"patient"`
	Provider *Provider `json:"provider"`
}

type NewAppointment struct {
	Date       string `json:"date"`
	PatientID  string `json:"patientId"`
	ProviderID string `json:"providerId"`
}

type NewPatient struct {
	Name string `json:"name"`
}

type NewProvider struct {
	Name string `json:"name"`
}

type Patient struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Provider struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
