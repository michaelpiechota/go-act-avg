package actions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddAction(t *testing.T) {
	var expected = map[string]ActionData{}

	// TEST CASE: Base Case - Happy Path
	t.Log("Base Case")
	actual := AddAction(`{"action":"run", "time":100.00}`)
	if actual != nil{
		t.Error("AddAction error for base case", actual)
	}
	// Expected data
	expected["run"] = ActionData{
		Average: 100,
		UnaryOpCounter: 1,
	}
	assert.Equal(t, expected, TempData)


	// TEST CASE: Base Case - Invalid Input
	t.Log("Base Case - Invalid Input")
	actual = AddAction(`{"action":"run", "time":asdf}`)
	if actual != nil{
		assert.EqualError(t, actual, "invalid character 'a' looking for beginning of value")
	}


	// TEST CASE: Invalid JSON String Input
	t.Log("Base Case - Invalid Input")
	actual = AddAction(`asdf`)
	if actual != nil{
		assert.EqualError(t, actual, "invalid character 'a' looking for beginning of value")
	}


	// TEST CASE: Type Differential
	t.Log("Type Differential")
	actual = AddAction(`{"action":"run", "time":100}`)
	if actual != nil{
		t.Error("AddAction error for Type Differential test case", actual)
	}
	// Expected data
	expected["jump"] = ActionData{
		Average: 100,
		UnaryOpCounter: 1,
	}
	assert.NotEqual(t, expected, TempData)
}

// Note: TempData (placeholder db) data persists through test suite.
// This is taken into account for the following tests.
func TestAddActionUpdateAverage(t *testing.T) {
	var expected = map[string]ActionData{}

	// TEST CASE: Base Case (single record) Counter for Average Calculation
	t.Log("Counter Base Case Test")
	actual := AddAction(`{"action":"jump", "time":100}`)
	if actual != nil{
		t.Error("Error with counter", actual)
	}
	// Expected data
	expected["jump"] = ActionData{
		Average: 100,
		UnaryOpCounter: 1,
	}
	assert.Equal(t, expected["jump"].UnaryOpCounter, TempData["jump"].UnaryOpCounter)


	// TEST CASE: Counter - Multiple Records for Average Calculation
	t.Log("Counter Multiple Records Test")
	actual = AddAction(`{"action":"jump", "time":100}`)
	if actual != nil{
		t.Error("Error with counter", actual)
	}
	// Expected data
	expected["jump"] = ActionData{
		Average: 100,
		UnaryOpCounter: 2,
	}
	assert.Equal(t, expected["jump"].UnaryOpCounter, TempData["jump"].UnaryOpCounter)


	//// TEST CASE: Average Calculation
	t.Log("Average Calculation Test")
	actual = AddAction(`{"action":"jump", "time":123.456}`)
	if actual != nil{
		t.Error("Error Calculating Average", actual)
	}
	// Expected data
	expected["jump"] = ActionData{
		Average: 107.81866666666667,
		UnaryOpCounter: 3,
	}
	assert.Equal(t, expected["jump"].Average, TempData["jump"].Average)


	//// TEST CASE: Negative input - Average Calculation
	t.Log("Negative input - Average Calculation Test")
	actual = AddAction(`{"action":"jump", "time":-123.456}`)
	if actual != nil{
		t.Error("Error Calculating Average with negative input", actual)
	}
	// Expected data
	expected["jump"] = ActionData{
		Average: 50,
		UnaryOpCounter: 3,
	}
	assert.Equal(t, expected["jump"].Average, TempData["jump"].Average)
}
