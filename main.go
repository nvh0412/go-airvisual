package main

import (
	"context"
	"fmt"

	"github.com/go-airvisual/airvisual"
)

func main() {
	client := airvisual.NewClient()

	data, _, err := client.Countries.ListCountries(context.Background())

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}
}
