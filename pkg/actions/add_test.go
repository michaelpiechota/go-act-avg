package actions
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
//// Table driven tests
//func TestAddAction(t *testing.T) {
//	tests := []struct {
//		name  string
//		input string
//		want  map[string]AverageData{}
//	}{
//		{name: "Serialized string to JSON struct", input: `{"action":"jump", "time":100}`, want: "nil",
//		//{name: "wrong sep", input: "a/b/c", want: []string{"a/b/c"}},
//		//{name: "no sep", input: "abc", want: []string{"abc"}},
//		//{name: "trailing sep", input: "a/b/c/", want: []string{"a", "b", "c"}},
//	}
//
//	for _, tc := range tests {
//		got := AddAction(tc.input, tc.sep)
//		if !reflect.DeepEqual(tc.want, got) {
//			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
//		}
//	}
//}
