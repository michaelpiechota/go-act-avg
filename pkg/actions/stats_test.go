package actions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStats(t *testing.T) {
	var expectedGetStats = map[string]ActionData{}

	// TEST CASE: No Data Test
	t.Log("No Data Case")
	actual := GetStats()
	if actual == ""{
		assert.Equal(t, "", actual)
	}


	// TEST CASE: Base Case - Happy Path
	t.Log("Base Case")

	// add some data
	TempData["run"] = ActionData{
		Average: 100,
		UnaryOpCounter: 1,
	}

	actual = GetStats()
	if actual == "" {
		t.Error("GetStats error for base case", actual)
	}

	assert.Equal(t, expectedGetStats, TempData)




}

