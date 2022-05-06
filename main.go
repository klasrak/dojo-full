package main

import "swapi/api"

func main() {
	api := api.New()

	if err := api.Run(); err != nil {
		panic(err)
	}
}
