package actions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//
//import (
//	"reflect"
//	"testing"
//)
//
////func TestAddAction(t *testing.T) {
////
////}
//
// Table driven tests
//func TestAddAction(t *testing.T) {
//	tests := []struct {
//		name  string
//		input string
//		want  string
//	}{
//		{name: "Valid Input", input: `{"action":"jump", "time":100}`, want: "nil"},
//		{name: "Invalid Input", input: `{"action:jump", "time":100}`, want: ""},
//		//{name: "no sep", input: "abc", want: []string{"abc"}},
//		//{name: "trailing sep", input: "a/b/c/", want: []string{"a", "b", "c"}},
//	}
//
//	for _, tc := range tests {
//		got := AddAction(tc.input)
//		if !reflect.DeepEqual(tc.want, got) {
//			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
//		}
//	}
//}

func TestAddAction(t *testing.T) {
	var expected = map[string]ActionData{}

	// TEST CASE: Base Case
	t.Log("Base Case")
	actual := AddAction(`{"action":"run", "time":100}`)
	if actual != nil{
		t.Error("AddAction error for base case", actual)
	}
	// Expected data
	expected["run"] = ActionData{
		Average: 100,
		UnaryOpCounter: 1,
	}
	assert.Equal(t, expected, TempData)


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

// Note: TempData data persists through test suite.
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


	// TEST CASE: Counter Multiple Records for Average Calculation
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
	//t.Log("Average Calculation Test")
	//actual = AddAction(`{"action":"jump", "time":777}`)
	//if actual != nil{
	//	t.Error("Error Calculating Average", actual)
	//}
	//// Expected data
	//expected["jump"] = ActionData{
	//	Average: 100,
	//	UnaryOpCounter: 3,
	//}
	//assert.Equal(t, expected["jump"].Average, TempData["jump"].Average)

}

