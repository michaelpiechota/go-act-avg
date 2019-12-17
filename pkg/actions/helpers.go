package actions

import (
	"go.uber.org/zap"
	"log"
)

// Helper functions

func getService() *Service {
	svc, err := NewService()
	if err != nil {
		log.Panicf("ERROR: unable to create service - %v", err)
	}
	return svc
}

// updateAverage checks record for an average based on action, updates as needed.
// If there is no average, the current action
// is the initial action; set average to current input time.
// If an average exists, calculate the new average and update record.
func updateAverage(input Input, svc Service) error {
	// grab latest, possibly empty
	average := TempData[input.Action]

	// check if an average exists
	if (ActionData{}) == average {
		initialAverage := &ActionData{
			Average:        float64(input.Time),
			UnaryOpCounter: 1,
		}
		// update record with initial average
		TempData[input.Action] = *initialAverage
		return nil
	} else {
		// Calculate new average
		na := average.Average + ((float64(input.Time) - average.Average) / float64(average.UnaryOpCounter+1))

		// create data object to update record
		updateAverage := &ActionData{
			Average:        na,
			UnaryOpCounter: average.UnaryOpCounter + 1,
		}

		// update record with updated average
		TempData[input.Action] = *updateAverage

		svc.logger.Info("Current data", zap.Any("TempData", TempData))

		return nil
	}
}
