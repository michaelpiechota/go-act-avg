package actions

import (
	"encoding/json"
	"log"
)



// AddAction accepts a json serialized string of the form:
// {"action":"string", "time":int}
// and maintains an average time for each action.
func AddAction(s string) error {
	// initialize input model
	input := Input{}
	// unmarshal serialized json into json struct
	err := json.Unmarshal([]byte(s), &input)
	if err != nil{
		log.Fatal("Failed to unmarshal input ", err)
	}





	return nil
}

