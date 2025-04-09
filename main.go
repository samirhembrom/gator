package main

import (
	"fmt"
	"log"

	"github.com/samirhembrom/blogaggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %v\n", cfg)

	err = cfg.SetUser("sam")

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %v\n", cfg)
}
