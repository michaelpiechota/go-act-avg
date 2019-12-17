package actions

import (
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"sync"
)

// AddAction accepts a json serialized string of the form:
// {"action":"string", "time":int}
// and maintains an average time for each action.
func AddAction(s string) error {
	svc := getService()
	// uncomment for granular input logging
	//svc.logger.Info("Action input received", zap.Any("Action", s))

	// unmarshal serialized json into input model
	err := json.Unmarshal([]byte(s), &svc.input)
	if err != nil {
		svc.logger.Error("failed to unmarshal input", zap.Error(err))
		return err
	}

	// TODO: Implementation docs for choosing to use mutex
	// Using a mutex to synchronize access to "TempData" map
	// during read/write ops.
	var mutex = sync.RWMutex{}
	mutex.Lock()

	error := updateAverage(svc.input, *svc)
	if error != nil {
		return errors.New("error calculating average")
	}

	mutex.Unlock()

	return nil
}
