package models

// Data model for AddAction input
// serialized JSON string -> JSON struct conversion
type Input struct {
	Action string	`json:"action"`
	Time int		`json:"time"`
}

