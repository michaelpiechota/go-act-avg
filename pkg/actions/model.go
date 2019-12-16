package actions

// Data model for AddAction input
// serialized JSON string -> JSON struct conversion
type Input struct {
	Action string	`json:"action"`
	Time int		`json:"time"`
}

// Data model for temp local data store.
// Average will hold the average time.
// Using float32
// UnaryOpCounter tracks running total needed
// for calculating average.
// TODO: big.Float or float64 reasoning
type ActionData struct{
	Average float64
	UnaryOpCounter int
}

type Stats struct {
	Action string	`json:"action"`
	Average float64	`json:"average"`
}
