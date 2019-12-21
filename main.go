package main

import (
	"context"
	"fmt"

	"github.com/go-airvisual/airvisual"
)

func main() {
	client := airvisual.NewClient()

	ctxBackground := context.Background()

	ctx2 := context.WithValue(ctxBackground, "API_KEY", "<API_KEY>")

	states, _, err := client.States.ListStates(ctx2, "Vietnam")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(states)
}
