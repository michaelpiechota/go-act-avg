package main

import (
	"fmt"
	"github.com/michaelpiechota/go-act-avg/pkg/actions"
)

func main() {
	actions.AddAction(`{"action":"jump", "time":100}`)
	getStats := actions.GetStats()
	fmt.Println(getStats)

	actions.AddAction(`{"action":"run", "time":75}`)
	getStats2 := actions.GetStats()
	fmt.Println(getStats2)

	actions.AddAction(`{"action":"jump", "time":200}`)
	getStats3 := actions.GetStats()
	fmt.Println(getStats3)
}
