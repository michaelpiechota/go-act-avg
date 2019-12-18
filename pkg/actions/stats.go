package actions

import (
	"encoding/json"
	"go.uber.org/zap"
)

func GetStats() string {
	svc := getService()

	// check if any data exists
	if len(TempData) == 0 {
		svc.logger.Info("no statistics available", zap.Any("TempData", TempData))
		return ""
	}

	// if data exists, fill model
	for action, average := range TempData {
		svc.stats = append(svc.stats, Stats{
			Action:  action,
			Average: average.Average,
		})
	}

	// serialize data
	serializedStats, err := json.Marshal(svc.stats)
	if err != nil {
		svc.logger.Error("failed to marshal input", zap.Error(err))
		return err.Error()
	}

	return string(serializedStats)
}
