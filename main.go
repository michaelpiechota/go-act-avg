package main

import (
	"fmt"
	"github.com/michaelpiechota/go-act-avg/pkg/pkg/actions"
)

func main (){
	actions.AddAction(`{"action": "run", "time": 350}`)

	fmt.Println("main works")

}
