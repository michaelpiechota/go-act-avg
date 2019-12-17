package actions

import (
	"go.uber.org/zap"
)

type Service struct {
	logger *zap.Logger
	input  Input
	stats  []Stats
}

func NewService() (*Service, error) {
	var err error

	// build custom logger configuration
	logger, _ := cfg.Build()
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
