package actions

import (
	"go.uber.org/zap"
	"log"
)

type Service struct {
	logger *zap.Logger
	input  Input
	stats  []Stats
}

func NewService() (*Service, error) {
	var err error

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
	}

	return &svc, err
}

func getService() *Service {
	svc, err := NewService()
	if err != nil {
		log.Panicf("ERROR: unable to create service - %v", err)
	}
	return svc
}
