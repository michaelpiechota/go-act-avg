package main

import (
	"github.com/michaelpiechota/go-act-avg/pkg/actions"
	"fmt"
)

func main() {
	actions.AddAction(`{"action":"jump", "time":100}`)
	actions.AddAction(`{"action":"run", "time":75}`)
	actions.AddAction(`{"action":"jump", "time":200}`)

	getStats := actions.GetStats()
	fmt.Println(getStats)
}
