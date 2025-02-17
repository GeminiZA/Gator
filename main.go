package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/GeminiZA/Gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	cfg.SetUser("tristan")
	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)
}
