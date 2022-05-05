package main

import "swapi/api"

func main() {
	if err := api.Run(); err != nil {
		panic(err)
	}
}
