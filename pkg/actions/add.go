package actions

import (
	"encoding/json"
	"errors"
	"log"
	"sync"
)

// AddAction accepts a json serialized string of the form:
// {"action":"string", "time":int}
// and maintains an average time for each action.
func AddAction(s string) error {
	// instantiate input model
	i := Input{}
	// unmarshal serialized json into input model
	err := json.Unmarshal([]byte(s), &i)
	if err != nil{
		log.Fatal("Failed to unmarshal input ", err)
	}

	// TODO: Implementation docs for choosing to use mutex
	var mutex = sync.RWMutex{}
	mutex.Lock()
	defer mutex.Unlock()

	error := updateAverage(i)
	if error != nil{
		return errors.New("error calculating average")
	}

	return nil
}





