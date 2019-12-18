package actions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStats(t *testing.T) {
	TempData = map[string]ActionData{}

	// TEST CASE: No Data Test
	t.Log("No Data Case")
	actual := GetStats()
	assert.Equal(t, "", actual)

	// TEST CASE: Base Case - Happy Path
	t.Log("Base Case")

	TempData["jump"] = ActionData{
		Average:        100,
		UnaryOpCounter: 1,
	}

	actual = GetStats()
	assert.Equal(t, `[{"action":"jump","average":100}]`, actual)

	// TEST CASE: Multiple stats
	t.Log("Multple Stats Test Case")

	TempData["jump"] = ActionData{
		Average:        100,
		UnaryOpCounter: 1,
	}
	TempData["run"] = ActionData{
		Average:        75,
		UnaryOpCounter: 1,
	}

	actual = GetStats()
	assert.Equal(t, `[{"action":"jump","average":100},{"action":"run","average":75}]`, actual)
}
