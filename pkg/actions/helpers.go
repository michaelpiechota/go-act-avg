package actions
// Helper functions

// updateAverage checks record for an average based on action, updates as needed.
// If there is no average, the current action
// is the initial action; set average to current input time.
// If an average exists, calculate the new average and update record.
func updateAverage(input Input) error {
	average := TempData[input.Action]

	// check if an average exists
	if (ActionData{}) == average {
		initialAverage := &ActionData{
			Average: float32(input.Time),
			UnaryOpCounter: 1,
		}

		// update record with initial average
		TempData[input.Action] = *initialAverage
		return nil
	}

	// Calculate new average
	na := average.Average + ((float32(input.Time) - average.Average) / float32(average.UnaryOpCounter + 1))

	// create data object to update record
	updateAverage := &ActionData{
		Average: na,
		UnaryOpCounter: average.UnaryOpCounter + 1,
	}

	// update record with updated average
	TempData[input.Action] = *updateAverage

	return nil
}
