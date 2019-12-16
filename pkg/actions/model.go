package actions

// Data model for AddAction input
// serialized JSON string -> JSON struct conversion
type Input struct {
	Action string	`json:"action"`
	Time int		`json:"time"`
}

// Data model for temp local data store.
// Average will hold the average time.
// UnaryOpCounter tracks running total needed
// for calculating average
// TODO: big.Float or float64 reasoning
type ActionData struct{
	Average float32
	UnaryOpCounter int
}

type Stats struct {
	Action string	`json:"action"`
	Average float32	`json:"average"`
}
