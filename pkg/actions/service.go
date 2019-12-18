package actions

import (
	"go.uber.org/zap"
	"sync"
)

type Service struct {
	logger *zap.Logger
	input  Input
	stats  []Stats
	mutex  sync.RWMutex
}

func NewService() (*Service, error) {
	var err error

	// mutex needed to sync state for data store
	// avoids concurrent write issues to map
	var m = sync.RWMutex{}

	// build custom logger configuration
	logger, error := cfg.Build()
	if error != nil {
		err = error
	}
	// flushes buffer, if any
	defer logger.Sync()

	// instantiate models
	i := Input{}
	var s []Stats

	svc := Service{
		logger: logger,
		input:  i,
		stats:  s,
		mutex:  m,
	}

	return &svc, err
}
